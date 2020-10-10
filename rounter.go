package main

import (
	"order-ops/controllers"
	"order-ops/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitGin(db *gorm.DB) *gin.Engine {
	orderService := services.NewOrderService()
	ctl := controllers.Controller{
		OrderService: orderService,
	}

	engine := gin.Default()
	engine.GET("/health", ctl.HealthCheck)
	apiGroup := engine.Group("/api/v1")
	{
		orderGroup := apiGroup.Group("/orders")
		{
			orderGroup.POST("", ctl.AddOrder)
		}
	}
	return engine
}
