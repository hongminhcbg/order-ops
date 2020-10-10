package dtos

type Meta struct {
	Code    int    `json:code`
	Message string `json:message`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
type Order struct {
	OrderNumber string `json:"orderNumber"`
	Name        string `json:"name"`
	Item        string `json:"item"`
	Quantiny    int32  `json:"quantiny"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}

type AddOrderRequest struct {
	Orders []Order `json:"orders"`
}

type AddorderResponse struct {
	ID int64 `json:"id"`
}

type LableDetails struct {
	TrackingNumber        string `json:"trackingNumber"`
	URL                   string `json:"url"`
	PartnerTrackingNumber string `json:"partnerTrackingNumber"`
}

type AddLabelRequest struct {
	OrderNumber  string       `json:"orderNumber"`
	LableDetails LableDetails `json:"lableDetails"`
}

type ShippingInfor struct {
	Status        int32  `json:"status"`
	BeginShipping string `json:"beginShipping"`
	TimeCompleted string `json:"timeCompleted"`
}
type AddShippingTimeRequest struct {
	OrderNumber   string `json:"orderNumber"`
	BeginShipping string `json:"beginShipping"`
	TimeCompleted string `json:"timeCompleted"`
}

type FullOrderInformation struct {
	AddOrderRequest
	ShippingInfor
	LableDetails LableDetails `json:"lableDetails"`
}

type SearchQuery struct {
	Key   string
	Value interface{}
}

type ChangeStatusToCompleted struct {
	OrderNumber string `json:"orderNumber"`
}