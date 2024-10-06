package httpserver

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type HTTPServer struct {
	gin  *gin.Engine
	port int
	db   *mongo.Database
}

func New(db *mongo.Database) *HTTPServer {
	return &HTTPServer{
		gin:  gin.Default(),
		port: 8001,
		db:   db,
	}
}
