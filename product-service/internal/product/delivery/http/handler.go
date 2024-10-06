package http

import (
	"net/http"

	model "github.com/Thehien0811/order-product-microservice/internal/models"
	"github.com/Thehien0811/order-product-microservice/internal/product/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderReq struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}

type UpdateProductInput struct {
	Quantity int32 `json:"quantity"`
}

func (h handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	input := usecase.CreateProductInput{
		Name:     req.Name,
		Quantity: req.Quantity,
	}

	o, err := h.uc.CreateProduct(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Order created successfully",
		"data":    o,
	})
}

func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	o, err := h.uc.DetailProduct(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   o,
	})
}

func (h handler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	var req UpdateProductInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	p, err := h.uc.DetailProduct(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get order",
			"details": err.Error(),
		})
		return
	}

	oid, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "failed",
		})
	}

	newP := model.Product{
		ID:       oid,
		Quantity: req.Quantity,
	}
	_, err = h.uc.UpdateProduct(ctx, newP)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "failed",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
	})
}
