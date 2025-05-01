package common

import (
	"context"
	"net/http"

	"github.com/ebhlz88/est-shop/models"
)

type ApiFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// common/interfaces.go

type Store interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
	AddProduct(name string, buyPrice int64, sellPrice int64) error
	ModifyProduct(id int, name string, buyPrice int64, sellPrice int64) error
	DeleteProduct(id int) error
}

type JSONApiServer interface {
	GetStore() Store
}
