package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

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
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
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
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("get raw body error", err)
		utils.ResponseErrorGin(ctx, "get raw body error")
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		fmt.Println("bind json error", err, "raw_body", string(bytes))
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
	result := make([]dtos.SearchQuery, 0)
	begin := ctx.Query("begin")
	if begin != "" {
		item := dtos.SearchQuery{
			Key:   "created_at > ?",
			Value: begin,
		}
		result = append(result, item)
	}

	end := ctx.Query("end")
	if end != "" {
		item := dtos.SearchQuery{
			Key:   "created_at < ?",
			Value: end,
		}
		result = append(result, item)
	}

	orderNumber := ctx.Query("order_number")
	if orderNumber != "" {
		item := dtos.SearchQuery{
			Key:   "order_number = ?",
			Value: orderNumber,
		}
		result = append(result, item)
	}

	return result, nil
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

	begin, err1 := time.Parse(services.CommonTimeFormat, request.BeginShipping)
	complete, err2 := time.Parse(services.CommonTimeFormat, request.TimeCompleted)
	if err1 != nil || err2 != nil {
		fmt.Println("time parser error", err1, err2)
		utils.ResponseErrorGin(ctx, "time parser error")
		return
	}

	request.BeginShippingReal = &begin
	request.TimeCompletedReal = &complete

	resp, err := c.OrderService.AddShippingTime(request)
	if err != nil {
		fmt.Println("add shipping time error", err)
		utils.ResponseErrorGin(ctx, "add shipping time error")
		return
	}

	fmt.Println("add shipping time success")
	utils.ResponseSuccess(ctx, resp)
}
