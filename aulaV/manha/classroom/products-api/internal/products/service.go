package products

import (
	"errors"

	"github.com/batatinha123/products-api/internal/entities"
)

var (
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotFound      = errors.New("product not found")
)

type Service interface {
	GetAll() ([]entities.Product, error)
	Store(name, category string, count int, price float64) (entities.Product, error)
	Update(id uint64, name, category string, count int, price float64) (entities.Product, error)
	UpdateName(id uint64, name string) (entities.Product, error)
	Delete(id uint64) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]entities.Product, error) {
	produtos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return produtos, nil

}

func (s *service) Store(name, category string, count int, price float64) (entities.Product, error) {
	// aqui poderiamos também através da service enviar o id ao repositório caso quisessemos
	// lastID, err := s.repository.LastID()
	// if err != nil {
	// 	return entities.Product{}, err
	// }

	// lastID++

	product, err := s.repository.Store(name, category, count, price)
	if err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (s *service) Update(id uint64, name, category string, count int, price float64) (entities.Product, error) {
	return s.repository.Update(id, name, category, count, price)
}

func (s *service) UpdateName(id uint64, name string) (entities.Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
