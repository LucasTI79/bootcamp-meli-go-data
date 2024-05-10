package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/batatinha123/products-api/cmd/server/handler"
	"github.com/batatinha123/products-api/internal/entities"
	"github.com/batatinha123/products-api/internal/products"
	"github.com/batatinha123/products-api/pkg/store"
	"github.com/batatinha123/products-api/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// condições

// usuário já deve ter sido criado
// usuário deve ter um token de autenticação
// deve haver produtos cadastrados

// cenário
// deve ser possivel buscar os produtos cadastrados da aplicação

func Test_GetProduct(t *testing.T) {
	t.Run("Test_GetProduct_OK", func(t *testing.T) {
		var response web.Response

		// criar um servidor e define suas rotas
		r := createServer()

		req, rr := createRequestTest(http.MethodPost, "/products/", `{
			"name": "Produto 1",
			"category": "teste",
			"count": 10,
			"price": 20.00
		}`)

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		// criar uma Request do tipo GET e Response para obter o resultado
		req, rr = createRequestTest(http.MethodGet, "/products/", "")

		// diz ao servidor que ele pode atender a solicitação
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Nil(t, err)
		var data []entities.Product
		jsonData, err := json.Marshal(response.Data)
		assert.Nil(t, err)
		err = json.Unmarshal(jsonData, &data)
		assert.Nil(t, err)
		assert.True(t, len(data) > 0)
	})

	t.Run("Test_GetProduct_No_Content", func(t *testing.T) {
		// criar um servidor e define suas rotas
		r := createServer()
		// criar uma Request do tipo GET e Response para obter o resultado
		req, rr := createRequestTest(http.MethodGet, "/products/", "")

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewFileStore(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api_token", "123456")

	return req, httptest.NewRecorder()
}
