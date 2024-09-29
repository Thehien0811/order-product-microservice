package httpserver

import "github.com/gin-gonic/gin"

type HTTPServer struct {
	gin  *gin.Engine
	port int
}
func New() *HTTPServer {
	return &HTTPServer{
		gin:          gin.Default(),
		port:         8001,
	}
}
