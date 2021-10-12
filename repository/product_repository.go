package repository

import "produtos/model"

type ProductRepository interface {
	Insert(product model.Product) model.Product

	FindAll() (products []model.Product)
}
