package http

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.RouterGroup, h Handler) {
	group := r.Group("/")
	group.POST("", h.Create)
	group.GET("/:id", h.Get)
}
