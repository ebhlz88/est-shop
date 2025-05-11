package common

import (
	"context"
	"net/http"
	"time"

	"github.com/ebhlz88/est-shop/models"
)

type ApiFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type Store interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
	AddProduct(name string, buyPrice int64, sellPrice int64) error
	ModifyProduct(id int, name string, buyPrice int64, sellPrice int64) error
	DeleteProduct(id int) error
	CreateOrder(productId int, amount int, isProfitDistributed bool) error
	GetAllOrders() ([]models.OrderWithProduct, error)
	CreateUser(name, username, password string, number int, createAt time.Time) error
	GetUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

type JSONApiServer interface {
	GetStore() Store
}
