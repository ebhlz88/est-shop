package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ebhlz88/est-shop/common"
	"github.com/ebhlz88/est-shop/models"
	"github.com/ebhlz88/est-shop/utils"
)

func HandleOrder(s common.Store) common.ApiFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if r.Method == http.MethodGet {
			return GetOrders(w, r, s)
		}
		if r.Method == http.MethodPost {
			return CreateOrder(w, r, s)
		}
		return nil
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request, s common.Store) error {
	var order models.Order

	// Decode the request body into an Order struct
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
			Error: err.Error(),
		})
	}

	// Pass only the ProductId (an int) to the CreateOrder function
	if err := s.CreateOrder(order.ProductId, order.Amount, order.IsProfitDistributed); err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err.Error(),
		})
	}

	return utils.WriteJson(w, http.StatusOK, models.APISuccessMessage{
		Message: "Successfully Added",
	})
}

func GetOrders(w http.ResponseWriter, r *http.Request, s common.Store) error {
	orders, err := s.GetAllOrders()
	fmt.Print(err)
	if err != nil {
		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
			Error: err.Error(),
		})
	}
	return utils.WriteJson(w, http.StatusOK, orders)
}
