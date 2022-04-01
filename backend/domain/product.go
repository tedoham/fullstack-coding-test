package domain

import (
	"time"

	"github.com/tedoham/fullstack-coding-test/dto"
	"github.com/tedoham/fullstack-coding-test/errs"
)

type Product struct {
	ID                    int       `db:"id"`
	ProductTypeID         int       `db:"product_type_id"`
	CustomerID            int       `db:"customer_id"`
	DeliveryStatus        int       `db:"delivery_status"`
	DeliveryAddress       string    `db:"delivery_address"`
	EstimatedDeliveryDate time.Time `db:"estimated_delivery_date"`
}

type ProductList struct {
	ProductId             string    `db:"id"`
	ProductTypeName       string    `db:"product_type_name"`
	CustomerName          string    `db:"customer_name"`
	DeliveryStatus        string    `db:"delivery_status"`
	DeliveryAddress       string    `db:"delivery_address"`
	EstimatedDeliveryDate time.Time `db:"estimated_delivery_date"`
}

func (p ProductList) ToDto() dto.ProductResponse {
	return dto.ProductResponse{
		ProductId:             p.ProductId,
		ProductTypeName:       p.ProductTypeName,
		CustomerName:          p.CustomerName,
		DeliveryStatus:        p.DeliveryStatus,
		DeliveryAddress:       p.DeliveryAddress,
		EstimatedDeliveryDate: p.EstimatedDeliveryDate,
	}
}

type ProductRepository interface {
	GetAllProducts(limit int, offset int) ([]ProductList, *errs.AppError)
	GetProductById(productId int) (*ProductList, *errs.AppError)
	UpdateDeliveryStatus(productId int, deliveryStatusId int) *errs.AppError
}
