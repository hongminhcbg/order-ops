package dtos

type Meta struct {
	Code    int    `json:code`
	Message string `json:message`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

type AddOrderRequest struct {
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

type AddorderResponse struct {
	ID int64 `json:"id"`
}
