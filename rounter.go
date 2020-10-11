package main

import (
	"order-ops/controllers"
	"order-ops/daos"
	"order-ops/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitGin(db *gorm.DB) *gin.Engine {
	orderDao := daos.NewOrderDao(db)
	orderService := services.NewOrderService(orderDao)

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
			orderGroup.DELETE("", ctl.Delete)
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
