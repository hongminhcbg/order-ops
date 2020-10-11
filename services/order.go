package services

import (
	"fmt"
	"order-ops/daos"
	"order-ops/dtos"
	"order-ops/models"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type OrderService interface {
	AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error)
	AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	Search(queries []dtos.SearchQuery) ([]dtos.FullOrderInformation, error)
	AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	Detete(orderNumber string) error
	Updates(request dtos.Order) (*dtos.AddorderResponse, error)
}

type orderServiceImpl struct {
	dao daos.OrderDao
}

func NewOrderService(dao daos.OrderDao) OrderService {
	return &orderServiceImpl{
		dao: dao,
	}
}

const CommonTimeFormat = "2006-01-02 15:04:05"
const (
	processingStatus = 0
	shippingStatus   = 1
	holdOnStatus     = 2
	completedStatus  = 3
)

func (service *orderServiceImpl) mapperDtossToModelOrder(input dtos.Order) models.Order {
	return models.Order{
		OrderNumber:  input.OrderNumber,
		CustomerName: input.Name,
		Quantity:     input.Quantity,
		Phone:        input.Phone,
		Address1:     input.Address1,
		Address2:     input.Address2,
		City:         input.City,
		State:        input.State,
		PostalCode:   input.PostalCode,
		Country:      input.Country,
		Note:         input.Note,
	}
}

func (service *orderServiceImpl) mapperDtossToModelOrderAddLable(input dtos.AddLabelRequest) models.Order {
	return models.Order{
		OrderNumber:           input.OrderNumber,
		TrackingNumber:        input.LableDetails.TrackingNumber,
		URL:                   input.LableDetails.URL,
		PartnerTrackingNumber: input.LableDetails.PartnerTrackingNumber,
	}
}

func (service *orderServiceImpl) AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, order := range request.Orders {
		record := service.mapperDtossToModelOrder(order)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, order.OrderNumber)
		} else {
			recordSuccess = append(recordSuccess, order.OrderNumber)
		}
	}

	result := dtos.AddorderResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *orderServiceImpl) AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error) {
	record := service.mapperDtossToModelOrderAddLable(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}

func (service *orderServiceImpl) mapperModelsToOrderFullInfor(input models.Order) dtos.FullOrderInformation {
	begin := input.BeginShipping.Format(CommonTimeFormat)
	end := input.TimeCompleted.Format(CommonTimeFormat)
	if begin == end {
		begin = ""
		end = ""
	}

	return dtos.FullOrderInformation{
		dtos.Order{
			OrderNumber: input.OrderNumber,
			Name:        input.CustomerName,
			Quantity:    input.Quantity,
			Phone:       input.Phone,
			Address1:    input.Address1,
			Address2:    input.Address2,
			City:        input.City,
			State:       input.State,
			PostalCode:  input.PostalCode,
			Country:     input.Country,
			Note:        input.Note,
		},
		dtos.ShippingInfor{
			Status:        input.Status,
			BeginShipping: begin,
			TimeCompleted: end,
		},
		dtos.LableDetails{
			TrackingNumber:        input.TrackingNumber,
			URL:                   input.URL,
			PartnerTrackingNumber: input.PartnerTrackingNumber,
		},
	}
}

func (service *orderServiceImpl) updateRecordState(input *models.Order) {
	if input.BeginShipping.Equal(*input.TimeCompleted) || input.Status == completedStatus {
		return
	}

	now := time.Now()
	if now.After(*input.BeginShipping) && now.Before(*input.TimeCompleted) {
		input.Status = shippingStatus
		return
	}

	if now.After(*input.BeginShipping) {
		input.Status = holdOnStatus
		return
	}
}

func (service *orderServiceImpl) Search(queries []dtos.SearchQuery) ([]dtos.FullOrderInformation, error) {
	records, _ := service.dao.Search(queries)
	result := make([]dtos.FullOrderInformation, 0)
	status := -1

	for _, query := range queries {
		if query.Key == "status=?" {
			statusint, _ := strconv.Atoi(fmt.Sprintf("%v", query.Value))
			status = statusint
		}
	}

	for _, record := range records {
		service.updateRecordState(&record)
		if status != -1 {
			if int(record.Status) == status {
				result = append(result, service.mapperModelsToOrderFullInfor(record))
			}
		} else {
			result = append(result, service.mapperModelsToOrderFullInfor(record))
		}
	}

	return result, nil
}

func (service *orderServiceImpl) AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error) {
	record := models.Order{
		OrderNumber:   request.OrderNumber,
		BeginShipping: request.BeginShippingReal,
		TimeCompleted: request.TimeCompletedReal,
	}
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}

func (service *orderServiceImpl) MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error) {
	record := models.Order{
		OrderNumber: orderNumber,
		Status:      completedStatus,
	}
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}

func (service *orderServiceImpl) Detete(orderNumber string) error {
	return service.dao.Delete(orderNumber)
}

func (service *orderServiceImpl) Updates(request dtos.Order) (*dtos.AddorderResponse, error) {
	record := service.mapperDtossToModelOrder(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, errors.Wrap(err, "update record error")
	}

	return &dtos.AddorderResponse{ID: record.ID}, nil
}
