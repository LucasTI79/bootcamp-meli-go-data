package products_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/batatinha123/products-api/internal/entities"
	"github.com/batatinha123/products-api/internal/products"
	"github.com/stretchr/testify/assert"
)

func Test_MySqlRepositoryWithSqlMock_Store_Mock(t *testing.T) {
	var productId uint64 = 1
	product := entities.Product{
		ID: productId,
	}

	db, mock := InitSqlMockDatabase(t)
	defer db.Close()

	smtMock := mock.
		ExpectPrepare(regexp.QuoteMeta(products.SaveProduct))

	smtMock.
		ExpectExec().
		WithArgs(product.Name, product.Category, product.Count, product.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	columns := []string{
		"id",
		"name",
		"type",
		"count",
		"price",
	}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "", "", "", "")

	mock.
		ExpectQuery("SELECT .* FROM products").
		WithArgs(productId).
		WillReturnRows(rows)

	repository := products.NewRepository(db)

	getResult, err := repository.GetOne(productId)
	assert.NoError(t, err)
	assert.Nil(t, getResult)

	productResult, err := repository.Store(product.Name, product.Category, product.Count, product.Price)
	assert.NoError(t, err)
	assert.NotNil(t, productResult)

	getResult, err = repository.GetOne(productId)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)

	assert.Equal(t, product.ID, getResult.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func InitSqlMockDatabase(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
