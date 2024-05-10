package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/batatinha123/products-api/internal/products"
	"github.com/batatinha123/products-api/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateRequestDto struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type UpdateRequestDto struct {
	Name     string  `json:"name" `
	Category string  `json:"category" binding:"required"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type UpdateNameRequestDto struct {
	Name string `json:"name"`
}

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(p) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body CreateRequestDto true "entities.Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		// quando chamamos a service, os dados já estarão tratados
		fmt.Println(req.Name, req.Category, req.Count, req.Price)
		p, err := c.service.Store(req.Name, req.Category, req.Count, req.Price)
		if err != nil {
			switch {
			case errors.Is(err, products.ErrProductAlreadyExists):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		ctx.JSON(http.StatusCreated, p)
	}
}

func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// forma de se fazermos uma conversão de alfa númerico para inteiro
		// strconv.Atoi(ctx.Param("id"))
		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var req UpdateRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(req.Name, req.Category, req.Count, req.Price)

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o nome do produto é obrigatório"})
			return
		}

		if req.Category == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "a categoria do produto é obrigatório"})
			return
		}

		if req.Count == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "a quantidade é obrigatória"})
			return
		}

		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O preço é obrigatório"})
			return
		}

		p, err := c.service.Update(id, req.Name, req.Category, req.Count, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// product, err = s.repo.FindByID(id)
		// if err != nil {
		//
		// }
		// if request.Name != "" {
		//	product.Name = request.Name
		// }
		//
		// s.repo.Update(id, product)

		// o paramâmetro 0 define para que o GO pegue os bits do processador em que está rodando a aplicação
		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var req UpdateNameRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}
		p, err := c.service.UpdateName(id, req.Name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		// poderiamos usar o http.StatusNoContent -> 204 tbm!
		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}
