package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context) {
	db := database.ConnectDB()
	orders := []models.Order{}
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", orders)
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    orders,
	})
}

func CreateOrder(ctx *gin.Context) {
	db := database.ConnectDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"msg":     "Failed in creating new order",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"msg":     "Order has been created",
	})

}

func UpdateOrder(ctx *gin.Context) {
	db := database.ConnectDB()
	order := models.Order{}
	success := true
	msg := ""
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	temp, _ := strconv.Atoi(ctx.Param("orderId"))
	order.ID = uint(temp)

	err := db.Where("id = ?", ctx.Param("orderId")).First(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "order_id is not found!",
		})
		return
	}

	for index, _ := range order.Items {
		if order.Items[index].LineItemId != 0 {
			order.Items[index].OrderId = order.ID
			order.Items[index].ID = order.Items[index].LineItemId
		} else {
			success = false
			msg = "line_item_id must be filled"
		}
	}

	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": success,
			"msg":     msg,
		})
		return
	}
	fmt.Printf("Value Update: %+v\n", order)
	fmt.Println("Order ID : ", ctx.Param("orderId"))
	err = db.Model(&order).Where("id = ?", ctx.Param("orderId")).Updates(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": success,
		"msg":     "Order has been updated",
	})

}

func DeleteOrder(ctx *gin.Context) {
	db := database.ConnectDB()
	order := models.Order{}

	err := db.Where("id = ?", ctx.Param("orderId")).First(&order).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "order_id is not found!",
		})
		return
	}

	err = db.Where("id = ?", ctx.Param("orderId")).Delete(&order).Error
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"msg":     "Order has been deleted",
	})
}
