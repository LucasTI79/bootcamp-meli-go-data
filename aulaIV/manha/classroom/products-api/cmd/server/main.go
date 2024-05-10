package main

import (
	"fmt"
	"log"
	"os"

	"github.com/batatinha123/products-api/cmd/server/handler"
	"github.com/batatinha123/products-api/cmd/server/middleware"
	"github.com/batatinha123/products-api/internal/products"
	"github.com/batatinha123/products-api/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/batatinha123/products-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// var dbConn *sql.DB

func GetDummyEndpoint(c *gin.Context) {
	resp := map[string]string{"hello": "world"}
	c.JSON(200, resp)
}

// middlewares globais
// middlewares de rota

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	user := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")

	fmt.Println("user", user, "pass", password)

	db := store.NewFileStore("file", "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// vai usar o middleware antes de cada handler
	server.Use(middleware.LoggerMiddleware)
	server.Use(middleware.DummyMiddleware())
	server.GET("/dummy", GetDummyEndpoint)
	pr := server.Group("/products")
	pr.Use(middleware.TokenAuthMiddleware())
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	pr.PUT("/:productId", productHandler.Update())
	pr.PATCH("/:productId", productHandler.UpdateName())
	pr.DELETE("/:productId", productHandler.Delete())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
