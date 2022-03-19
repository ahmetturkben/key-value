package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"record/domain"
	_file "record/key/repository/file"
	_inmemory "record/key/repository/inmemory"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "record/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_configurationHandler "record/key/delivery/http"
	_configurationRouter "record/key/delivery/router"
	_configurationUsecase "record/key/usecase"
)

func main() {

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	inMemDb, err := _inmemory.Init()
	if err != nil {
		log.Fatalf("InMemory Init method failed: %v", err)
		panic(err)
	}

	fileRepo := _file.NewFileRepository()
	memoryRepo := _inmemory.NewMemoryRepository(inMemDb)
	memoryUsecase := _configurationUsecase.NewMemoryUsecase(memoryRepo, fileRepo)
	memoryHandler := _configurationHandler.NewMemoryHandler(context.TODO(), memoryUsecase)

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/api/record/healthcheck", HealthCheck)
	_configurationRouter.MemoryRouter(app, memoryHandler)
	
	syncFile(memoryRepo, fileRepo)
	//time.Sleep(time.Second * 0)

	//port := ":" + os.Getenv("PORT")
	if err := app.Listen(":80"); err != nil {
		log.Fatal(err)
	}
}

func syncFile(repository domain.MemoryRepository, fileRepo domain.FileRepository) {

	readData := fileRepo.Read()
	var obj []*domain.KeyValue
	err := json.Unmarshal([]byte(readData), &obj)
	if err != nil {
		fmt.Println(err)
	}

	err = repository.CreateBulk(obj)
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup
	go func() {
		for range time.Tick(time.Second * 1) {
			wg.Add(1)
			fmt.Printf("Worker %d starting\n", time.Second)

			data, err := repository.GetAll()
			if err != nil {
				return
			}
			dbJson, _ := json.Marshal(data)
			fileRepo.Write(dbJson)

			wg.Done()
			fmt.Printf("Worker %d done\n", time.Second)
		}
	}()

	wg.Wait()
	fmt.Printf("Worker %d wait\n", time.Second)
}

func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
