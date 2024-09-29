package main

import (
	"github.com/Thehien0811/order-product-microservice/internal/httpserver"
)

func main() {
	srv := httpserver.New()
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
