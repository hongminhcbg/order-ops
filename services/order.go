package services

import (
	"order-ops/dtos"

	"github.com/pkg/errors"
)

type OrderService interface {
	AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error)
}

type orderServiceImpl struct {
}

func NewOrderService() OrderService {
	return &orderServiceImpl{}
}

func (service *orderServiceImpl) AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error) {
	return nil, errors.New("not implement")
}
