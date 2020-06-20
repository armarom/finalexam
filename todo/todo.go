package todo

import (
	"github.com/armarom/finalexam/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AuthMiddleware)

	router.POST("/customers", createCustomersHandler)
	router.GET("/customers/:id", getCustomersByIdHandler)
	router.GET("/customers", getCustomersHandler)
	router.PUT("/customers/:id", updateCustomersHandler)
	router.DELETE("/customers/:id", deleteCustomersHandler)

	return router
}
