package products_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/batatinha123/products-api/internal/entities"
	"github.com/batatinha123/products-api/internal/products"
	test_utils "github.com/batatinha123/products-api/tests/utils"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func Test_MySqlRepositoryWithTxDb_Store_Mock(t *testing.T) {
	product := entities.Product{}

	db, err := test_utils.InitTxDbDatabase(t)
	assert.NoError(t, err)
	repository := products.NewRepository(db)

	productResult, err := repository.Store(product.Name, product.Category, product.Count, product.Price)
	assert.NoError(t, err)
	assert.NotNil(t, productResult)

	getResult, err := repository.GetOne(productResult.ID)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
}

func Test_MySqlRepositoryWithTxDb_GetOneWithContext(t *testing.T) {
	product := entities.Product{
		Name: "test",
	}

	db, err := test_utils.InitTxDbDatabase(t)
	assert.NoError(t, err)
	repository := products.NewRepository(db)

	productResult, err := repository.Store(product.Name, product.Category, product.Count, product.Price)
	assert.NoError(t, err)
	assert.NotNil(t, productResult)

	// cria um context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	getResult, err := repository.GetOneWithContext(ctx, productResult.ID)
	fmt.Println(err)
	assert.Equal(t, product.Name, getResult.Name)
}
