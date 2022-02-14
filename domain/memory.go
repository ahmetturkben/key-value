package domain

import (
	"context"
	_ "record/key/delivery/model"
)

type KeyValue struct {
	Key   string
	Value string
}

type MemoryKeyReaderRepository interface {
	GetByKey(context.Context, string) (*KeyValue, error)
	GetAll() ([]*KeyValue, error)
}

type MemoryKeyWriterRepository interface {
	Create(context.Context, *KeyValue) error
	CreateBulk([]*KeyValue) error
	Delete(context.Context) error
}

type MemoryRepository interface {
	MemoryKeyReaderRepository
	MemoryKeyWriterRepository
}

type MemoryUsecase interface {
	GetByKey(context.Context, string) (*KeyValue, error)
	Delete(context.Context) error
	GetAll() ([]*KeyValue, error)
	Create(context.Context, *KeyValue) error
}
