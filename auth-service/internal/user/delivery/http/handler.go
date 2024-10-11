package http

import (
	"net/http"

	"github.com/Thehien0811/order-product-microservice/internal/user/usecase"
	"github.com/gin-gonic/gin"
	
)

type SignUpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var req SignInReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	input := usecase.SignInInput{
		Username: req.Username,
		Password: req.Password,
	}

	o, err := h.uc.SignIn(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Signin successfully",
		"data":    o,
	})
}

func (h handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var req SignInReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	input := usecase.SignUpInput{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := h.uc.SignUp(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Signup successful",
		"token":    token,
	})
}
