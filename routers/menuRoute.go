package routers

import (
	"github.com/codepnw/restaurant-api/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controllers.GetMenus())
	incomingRoutes.GET("/menu/:menus_id", controllers.GetMenu())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
