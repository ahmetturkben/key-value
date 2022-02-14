package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"record/domain"
	"record/key/delivery/model"
)

type MemoryHandler struct {
	MemoryUsecase domain.MemoryUsecase
	Context       context.Context
}

func NewMemoryHandler(context context.Context, memoryUsecase domain.MemoryUsecase) *MemoryHandler {
	return &MemoryHandler{MemoryUsecase: memoryUsecase, Context: context}
}

func (handler *MemoryHandler) GetMemoryByKey(fiberContext *fiber.Ctx) error {
	key := fiberContext.Params("key")

	record, err := handler.MemoryUsecase.GetByKey(handler.Context, key)

	if err != nil {
		fiberContext.JSON(generateResponse(500, "DB Connection exception"))
		return nil
	}

	fiberContext.JSON(record)
	return nil
}

func (handler *MemoryHandler) Flush(fiberContext *fiber.Ctx) error {

	err := handler.MemoryUsecase.Delete(handler.Context)

	if err != nil {
		fiberContext.JSON(generateResponse(500, "DB Connection exception"))
		return nil
	}

	fiberContext.JSON(generateResponse(200, "Flush Success"))
	return nil
}

func (handler *MemoryHandler) CreateKey(fiberContext *fiber.Ctx) error {
	request := new(domain.KeyValue)
	if err := fiberContext.BodyParser(request); err != nil {
		fiberContext.JSON(generateResponse(500, "Json parser exception"))
		return err
	}
	err := handler.MemoryUsecase.Create(handler.Context, request)

	if err != nil {
		fiberContext.JSON(generateResponse(500, "DB Connection exception"))
		return nil
	}

	fiberContext.JSON("")
	return nil
}

func generateResponse(code int, message string) model.Response {
	response := model.Response{
		Code: code,
		Msg:  message,
	}
	return response
}
