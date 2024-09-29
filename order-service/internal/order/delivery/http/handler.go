package http

import (
	"net/http"

	"github.com/Thehien0811/order-product-microservice/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

type CreateOrderReq struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
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

	input := usecase.CreateOrderInput{
		ID:       req.ID,
		Name:     req.Name,
		Quantity: req.Quantity,
	}

	o, err := h.uc.CreateOrder(ctx, input)
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
	o, err := h.uc.DetailOrder(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    o,
	})
}
