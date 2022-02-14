package usecase

import (
	"context"
	_ "fmt"
	"github.com/pkg/errors"
	"record/domain"
	_ "record/key/delivery/model"
	"record/pkg/httpErrors"
	"record/pkg/utils"
	_ "strings"
)

type memoryUsecase struct {
	memoryRepo domain.MemoryRepository
	fileRepo   domain.FileRepository
}

func NewMemoryUsecase(memoryRepo domain.MemoryRepository, fileRepo domain.FileRepository) *memoryUsecase {
	return &memoryUsecase{memoryRepo: memoryRepo, fileRepo: fileRepo}
}

func (usecase *memoryUsecase) GetByKey(ctx context.Context, key string) (*domain.KeyValue, error) {
	var config domain.KeyValue
	record, err := usecase.memoryRepo.GetByKey(ctx, key)
	if err != nil {
		return &config, err
	}
	return record, nil
}

func (usecase *memoryUsecase) Delete(ctx context.Context) error {
	err := usecase.memoryRepo.Delete(ctx)
	if err != nil {
		return err
	}

	usecase.fileRepo.Delete()

	return nil
}

func (usecase *memoryUsecase) GetAll() ([]*domain.KeyValue, error) {
	var config []*domain.KeyValue
	records, err := usecase.memoryRepo.GetAll()
	if err != nil {
		return config, err
	}
	return records, nil
}

func (usecase *memoryUsecase) Create(ctx context.Context, key *domain.KeyValue) error {
	if err := utils.ValidateStruct(ctx, key); err != nil {
		return httpErrors.NewBadRequestError(errors.WithMessage(err, "records.Create.ValidateStruct"))
	}
	err := usecase.memoryRepo.Create(ctx, &domain.KeyValue{
		Key:   key.Key,
		Value: key.Value,
	})
	if err != nil {
		return httpErrors.NewInternalServerError(errors.WithMessage(err, "records.Create"))
	}
	return nil
}
