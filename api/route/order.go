package route

import (
	"DBOTechnicalTest-Golang/api/controller"
	"DBOTechnicalTest-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func Order(r *gin.RouterGroup, c controller.Controller) {
	orderAPI := r.Group("/order")
	orderAPI.Use(middleware.VerifyToken())
	orderAPI.POST("", c.OrderController.Create)
	orderAPI.GET("/all", c.OrderController.GetAll)
	orderAPI.GET("/:id", c.OrderController.GetByID)
	orderAPI.PUT("", c.OrderController.Update)
	orderAPI.DELETE("/:id", c.OrderController.Delete)
}
