package service

import (
	"produtos/model"
	"produtos/repository"
)

type CreateProductService struct {
	ProductRepository repository.ProductRepository
}

func NewCreateProductService(productRepository *repository.ProductRepository) CreateProductService {
	return CreateProductService{
		ProductRepository: *productRepository,
	}
}

func (service *CreateProductService) Execute() (response model.Response) {
	product := model.Product{
		Id:       "1",
		Name:     "Produto",
		Price:    10,
		Quantity: 2,
	}

	service.ProductRepository.Insert(product)

	response = model.Response{
		Code:   200,
		Status: "OK",
		Data:   product,
	}

	return response
}
