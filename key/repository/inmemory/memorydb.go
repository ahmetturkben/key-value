package inmemory

import (
	"context"
	"github.com/hashicorp/go-memdb"
	"record/domain"
)

type memoryRepository struct {
	repo *memdb.MemDB
}

func NewMemoryRepository(repo *memdb.MemDB) *memoryRepository {
	return &memoryRepository{repo: repo}
}

func Init() (*memdb.MemDB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"keyvalue": &memdb.TableSchema{
				Name: "keyvalue",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Key"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return db, err
	}
	return db, err
}

func (repo *memoryRepository) Create(ctx context.Context, pair *domain.KeyValue) error {
	transact := repo.repo.Txn(true)
	if err := transact.Insert("keyvalue", pair); err != nil {
		transact.Abort()
		return err
	}
	transact.Commit()
	return nil
}

func (repo *memoryRepository) CreateBulk(keys []*domain.KeyValue) error {
	transact := repo.repo.Txn(true)
	for _, p := range keys {
		if err := transact.Insert("keyvalue", p); err != nil {
			return err
		}
	}
	transact.Commit()
	return nil
}

func (repo *memoryRepository) GetAll() ([]*domain.KeyValue, error) {
	transact := repo.repo.Txn(false)
	defer transact.Abort()

	var keys []*domain.KeyValue

	itrx, err := transact.Get("keyvalue", "id")
	if err != nil {
		return nil, err
	} else if itrx == nil {
		return nil, nil
	}
	for obj := itrx.Next(); obj != nil; obj = itrx.Next() {
		p := obj.(*domain.KeyValue)
		keys = append(keys, p)
	}

	return keys, nil
}

func (repo *memoryRepository) GetByKey(ctx context.Context, key string) (*domain.KeyValue, error) {
	transact := repo.repo.Txn(false)
	defer transact.Abort()

	itrx, err := transact.First("keyvalue", "id", key)
	if err != nil {
		return nil, err
	} else if itrx == nil {
		return nil, nil
	}
	return itrx.(*domain.KeyValue), nil
}

func (repo *memoryRepository) Delete(ctx context.Context) error {
	transact := repo.repo.Txn(true)
	defer transact.Abort()

	itrx, err := transact.DeleteAll("keyvalue", "id")
	if err != nil {
		return err
	} else if itrx == 0 {
		return nil
	}

	transact.Commit()
	return nil
}
