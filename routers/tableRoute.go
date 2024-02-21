package routers

import (
	"github.com/codepnw/restaurant-api/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
