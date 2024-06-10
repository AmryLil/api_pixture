package routes

import (
	"api/config"
	"api/handler"
	"api/middlewares"
	"api/repositories"
	"api/services"

	"github.com/gin-gonic/gin"
)

func UserDetailsRouters(c *gin.RouterGroup) {
	userDetailsRepo := repositories.NewUserDetailsRepo(config.DB)
	userDeatilsService := services.NewUserDetailsService(userDetailsRepo)
	userDetailsHandler := handler.NewUserDetailsHandler(userDeatilsService)

	user := c.Group("/user")
	user.POST("/adduserdetail", userDetailsHandler.AddUserDetails)
	user.Use(middlewares.JWTMiddleware())
	user.GET("/userdetails", userDetailsHandler.GetDataUser)
}
