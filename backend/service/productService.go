package service

import (
	"github.com/tedoham/fullstack-coding-test/domain"
	"github.com/tedoham/fullstack-coding-test/dto"
	"github.com/tedoham/fullstack-coding-test/errs"
)

type ProductService interface {
	GetProducts(limit int, offset int) ([]dto.ProductResponse, *errs.AppError)
	GetProduct(productId int) (*dto.ProductResponse, *errs.AppError)
	UpdateDeliveryStatus(productId int, deliveryStatusId int) *errs.AppError
}

type DefaultProductService struct {
	repo domain.ProductRepository
}

func (s DefaultProductService) GetProducts(limit int, offset int) ([]dto.ProductResponse, *errs.AppError) {
	products, err := s.repo.GetAllProducts(limit, offset)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ProductResponse, 0)
	for _, c := range products {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultProductService) GetProduct(id int) (*dto.ProductResponse, *errs.AppError) {
	c, err := s.repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func (s DefaultProductService) UpdateDeliveryStatus(productId int, deliveryStatusId int) *errs.AppError {
	err := s.repo.UpdateDeliveryStatus(productId, deliveryStatusId)
	if err != nil {
		return err
	}

	return err
}

func NewProductService(repo domain.ProductRepository) DefaultProductService {
	return DefaultProductService{repo}
}
