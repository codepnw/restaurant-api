package main

import (
	"os"

	"github.com/codepnw/restaurant-api/middleware"
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
	router.Use(middleware.Authentication())

	r.FoodRoutes(router)
	r.MenuRoutes(router)

	router.Run(":" + port)
}
