package services

import (
	"order-ops/dtos"

	"github.com/pkg/errors"
)

type OrderService interface {
	AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error)
	AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	Search(queries []dtos.SearchQuery) ([]dtos.FullOrderInformation, error)
	AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.FullOrderInformation, error)
	MakeCompleted(orderNumber string) (*dtos.FullOrderInformation, error)
}

type orderServiceImpl struct {
}

func NewOrderService() OrderService {
	return &orderServiceImpl{}
}

func (service *orderServiceImpl) AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error) {
	return nil, errors.New("not implement")
}

func (service *orderServiceImpl) AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error) {
	return nil, errors.New("not implement")
}

func (service *orderServiceImpl) Search(queries []dtos.SearchQuery) ([]dtos.FullOrderInformation, error) {
	return nil, errors.New("not implement")
}

func (service *orderServiceImpl) AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.FullOrderInformation, error) {
	return nil, errors.New("not implement")
}

func (service *orderServiceImpl) MakeCompleted(orderNumber string) (*dtos.FullOrderInformation, error) {
	return nil, errors.New("not implement")
}
