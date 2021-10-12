package repository

import (
	"produtos/model"
)

type productRepositoryImpl struct {
	Products []model.Product
}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{
		Products: []model.Product{},
	}
}

func (repository *productRepositoryImpl) Insert(product model.Product) model.Product {
	repository.Products = append(repository.Products, model.Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	})

	return product
}

func (repository *productRepositoryImpl) FindAll() (products []model.Product) {
	return repository.Products
}
