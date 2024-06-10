package middlewares

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(200)
				return
		}
		c.Next()
}}