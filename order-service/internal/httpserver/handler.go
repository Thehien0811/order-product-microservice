package httpserver

func (srv HTTPServer) mapHandlers() error {
	api := srv.gin.Group("/api/v1")

	orderHTTP.MapRoutes(api.Group("/settings"), settingsH)
}