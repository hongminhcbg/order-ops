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

func (c Controller) AddLabelToOrder(ctx *gin.Context) {
	var request dtos.AddLabelRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.AddLabelsToOrder(request)
	if err != nil {
		fmt.Println("add labels to order error", err)
		utils.ResponseErrorGin(ctx, "add labels to order error")
		return
	}

	fmt.Println("add labels to order success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) getSearchQuery(ctx *gin.Context) ([]dtos.SearchQuery, error) {
	return nil, nil
}

func (c Controller) Search(ctx *gin.Context) {
	queries, err := c.getSearchQuery(ctx)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.Search(queries)
	if err != nil {
		fmt.Println("search orders error", err)
		utils.ResponseErrorGin(ctx, "search order error")
		return
	}

	fmt.Println("search success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) MakeDone(ctx *gin.Context) {
	var request dtos.ChangeStatusToCompleted
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.MakeCompleted(request.OrderNumber)
	if err != nil {
		fmt.Println("add labels to order error", err)
		utils.ResponseErrorGin(ctx, "add labels to order error")
		return
	}

	fmt.Println("add labels to order success")
	utils.ResponseSuccess(ctx, resp)
}

func (c Controller) AddShippingTime(ctx *gin.Context) {
	var request dtos.AddShippingTimeRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("bind json error", err)
		utils.ResponseErrorGin(ctx, "bind json error")
		return
	}

	resp, err := c.OrderService.AddShippingTime(request)
	if err != nil {
		fmt.Println("add shipping time error", err)
		utils.ResponseErrorGin(ctx, "add shipping time error")
		return
	}

	fmt.Println("add shipping time success")
	utils.ResponseSuccess(ctx, resp)
}
