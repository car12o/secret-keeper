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

func main() {
	r := gin.Default()

	r.GET("/metrics", metrics.Get)

	s := r.Group("/secret", metrics.Middleware())
	{
		s.POST("", secret.Post)

		s.GET("/:hash", secret.Get)
	}

	r.Run()
}
