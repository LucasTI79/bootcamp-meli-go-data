package products

import (
	"github.com/batatinha123/products-api/internal/entities"
	"github.com/batatinha123/products-api/pkg/store"
)

type Repository interface {
	GetAll() ([]entities.Product, error)
	Store(name, category string, count int, price float64) (entities.Product, error)
	Update(id uint64, name, productType string, count int, price float64) (entities.Product, error)
	UpdateName(id uint64, name string) (entities.Product, error)
	Delete(id uint64) error
}

func NewRepository(db store.Store) Repository {
	return &FileRepository{db}
}
