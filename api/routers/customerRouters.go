package routers

import (
	"interview/api/controllers"

	"github.com/gin-gonic/gin"
)

func LoadCustomerPaths(r *gin.Engine) {
	r.POST("/NutriNET/Cliente", controllers.CreateCustomer)
	r.GET("/NutriNET/Cliente", controllers.FindCustomers)
	r.PUT("/NutriNET/Cliente", controllers.UpdateCustomer)
}