package productservice

import (
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/domain"
)

func ProductEntityToProtobuf(product *domain.Products) *Product {
	return &Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Unit:        product.Unit,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Unix(),
		UpdatedAt:   product.UpdatedAt.Unix(),
	}
}
