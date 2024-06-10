package main

import (
	"api/config"
	"api/middlewares"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.DBConnetcion()

	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	api := router.Group("/api")
	routes.UserRoutes(api)
	routes.UserDetailsRouters(api)
	router.Run()

}
