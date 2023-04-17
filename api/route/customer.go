package route

import (
	"DBOTechnicalTest-Golang/api/controller"
	"DBOTechnicalTest-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func Customer(r *gin.RouterGroup, c controller.Controller) {
	customerAPI := r.Group("/customer")
	customerAPI.POST("", c.CustomerController.Create)
	customerAPI.Use(middleware.VerifyToken())
	customerAPI.GET("/all", c.CustomerController.GetAll)
	customerAPI.GET("/:id", c.CustomerController.GetByID)
	customerAPI.PUT("", c.CustomerController.Update)
	customerAPI.DELETE("/:id", c.CustomerController.Delete)
}
