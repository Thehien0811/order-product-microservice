package httpserver

import (
	"fmt"
)

func (srv HTTPServer) Run() error {
	err := srv.mapHandlers()
	if err != nil {
		return err
	}
	srv.gin.Run(fmt.Sprintf(":%d", srv.port))

	return nil
}
