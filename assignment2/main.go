package main

import (
	"assignment2/controllers"
	"assignment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()

	router.GET("/orders", controllers.GetOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	router.Run(":4444")
}
