package http

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.RouterGroup, h Handler) {
	group := r.Group("/")
	group.POST("/signin", h.SignIn)
	group.POST("/signup", h.SignUp)
}
