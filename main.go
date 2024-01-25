package main

import (
	"os"

	r "github.com/codepnw/restaurant-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	r.UserRoutes(router)

	router.Run(":" + port)
}
