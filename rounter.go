package main

import (
	"order-ops/controllers"
	"order-ops/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CORSMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func InitGin(db *gorm.DB) *gin.Engine {
	orderService := services.NewOrderService()
	ctl := controllers.Controller{
		OrderService: orderService,
	}

	engine := gin.Default()
	engine.Use(CORSMiddleWare())

	engine.GET("/health", ctl.HealthCheck)
	apiGroup := engine.Group("/api/v1")
	{
		orderGroup := apiGroup.Group("/orders")
		{
			orderGroup.POST("", ctl.AddOrder)
			orderGroup.GET("/search", ctl.Search)
			orderGroup.POST("/make-done", ctl.MakeDone)
			orderGroup.POST("/shipping-time", ctl.AddShippingTime)
		}

		labelGroup := apiGroup.Group("/labels")
		{
			labelGroup.POST("", ctl.AddLabelToOrder)
		}
	}
	return engine
}
