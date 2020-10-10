package controllers

import (
	"fmt"

	"order-ops/dtos"
	"order-ops/services"

	"order-ops/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	OrderService services.OrderService
}

func (c Controller) HealthCheck(contex *gin.Context) {
	contex.JSON(200, gin.H{
		"status": "running",
	})
}

func (c Controller) AddOrder(ctx *gin.Context) {
	var request dtos.AddOrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.AddOrder(request)
	if err != nil {
		fmt.Println("add order error", err)
		utils.ResponseErrorGin(ctx, "add order error")
		return
	}

	fmt.Println("add success")
	utils.ResponseSuccess(ctx, resp)
}
