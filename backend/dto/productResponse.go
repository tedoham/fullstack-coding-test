package dto

import "time"

type ProductResponse struct {
	ProductId             string    `json:"product_id"`
	ProductTypeName       string    `json:"product_type_name"`
	CustomerName          string    `json:"customer_name"`
	DeliveryStatus        string    `json:"delivery_status"`
	DeliveryAddress       string    `json:"delivery_address"`
	EstimatedDeliveryDate time.Time `json:"estimated_delivery_date"`
}
