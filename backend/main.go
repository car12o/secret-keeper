package main

import (
	"fmt"

	"github.com/car12o/secret-keeper/database"
	"github.com/car12o/secret-keeper/metrics"
	"github.com/car12o/secret-keeper/secret"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := database.Init(); err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %s", err.Error()))
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/metrics", metrics.Get)

	s := r.Group("/secret", metrics.Middleware())
	{
		s.POST("", secret.Post)

		s.GET("/:hash", secret.Get)
	}

	r.Run()
}
