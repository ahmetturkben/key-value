package file

import (
	"io/ioutil"
	"log"
	"os"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

var (
	newFile *os.File
	err     error
)

func (repo *fileRepository) Create() {
	_, err := os.Stat("data.txt")
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err = os.Create("data.txt")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (repo *fileRepository) Delete() {
	_, err := os.Stat("data.txt")
	if err != nil {
		if !os.IsNotExist(err) {
			os.Remove("data.txt")
		}
	}
}

func (repo *fileRepository) Write(data []byte) {
	err = ioutil.WriteFile("data.txt", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *fileRepository) Read() []byte {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	return data
}
