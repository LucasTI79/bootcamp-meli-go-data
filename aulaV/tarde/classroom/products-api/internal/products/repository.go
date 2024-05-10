package products

import (
	"context"
	"database/sql"

	"github.com/batatinha123/products-api/internal/entities"
)

type Repository interface {
	GetAll() ([]entities.Product, error)
	GetOne(id uint64) (entities.Product, error)
	Store(name, category string, count int, price float64) (entities.Product, error)
	Update(id uint64, name, productType string, count int, price float64) (entities.Product, error)
	UpdateName(id uint64, name string) (entities.Product, error)
	Delete(id uint64) error
	GetFullData(id uint64) (entities.ProductFullData, error)
	GetOneWithContext(ctx context.Context, id uint64) (entities.Product, error)
}

func NewRepository(db *sql.DB) Repository {
	return &MySqlRepository{db}
}
