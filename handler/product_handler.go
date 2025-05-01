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

func HandleProduct(s common.Store) common.ApiFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if r.Method == "GET" {
			return HandleGetProducts(w, r, s)
		}
		if r.Method == "POST" {
			return HandlePostProducts(w, r, s)
		}
		return fmt.Errorf("%s Method not allowd", r.Method)
	}
}

func HandleProductById(s common.Store) common.ApiFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		id, err := utils.GetId(r)
		if err != nil {
			return err
		}
		if r.Method == "GET" {
			return HandleGetProductById(w, r, s, id)
		}

		if r.Method == "DELETE" {
			return HandleDeleteProduct(w, r, s, id)
		}
		if r.Method == http.MethodPatch {
			return handlePatchProduct(w, r, s, id)
		}
		return fmt.Errorf("%s Method not allowd", r.Method)
	}
}

func handlePatchProduct(w http.ResponseWriter, r *http.Request, s common.Store, id int) error {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err,
		})
	}
	if err := s.ModifyProduct(id, product.ProductName, product.ProductBuyPrice, product.ProductSellPrice); err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err,
		})
	}
	return utils.WriteJson(w, http.StatusOK, product)
}
func HandleGetProductById(w http.ResponseWriter, r *http.Request, s common.Store, id int) error {
	product, err := s.GetProductById(id)
	if err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}
	return utils.WriteJson(w, http.StatusOK, product)
}

func HandleGetProducts(w http.ResponseWriter, r *http.Request, s common.Store) error {
	products, err := s.GetAllProducts()
	if err != nil {
		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
			Error: err,
		})
	}
	return utils.WriteJson(w, http.StatusAccepted, products)
}

func HandlePostProducts(w http.ResponseWriter, r *http.Request, s common.Store) error {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
			Error: err,
		})
	}
	err := s.AddProduct(product.ProductName, product.ProductBuyPrice, product.ProductSellPrice)
	if err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err,
		})
	}
	return utils.WriteJson(w, http.StatusCreated, map[string]string{
		"ok": "product successfully added",
	})
}

func HandleDeleteProduct(w http.ResponseWriter, r *http.Request, s common.Store, id int) error {
	err := s.DeleteProduct(id)
	if err != nil {
		return utils.WriteJson(w, http.StatusInternalServerError, models.APIError{
			Error: err,
		})
	}
	return utils.WriteJson(w, http.StatusNoContent, models.APISuccessMessage{
		Message: "Delete Successful",
	})
}

// func HandlePriceFetch(ctx context.Context, w http.ResponseWriter, r *http.Request, s *server.JSONApiServer) error {
// 	ticker := r.URL.Query().Get("ticker")
// 	price, err := s.Svc.FetchPrice(ctx, ticker)
// 	if err != nil {
// 		return utils.WriteJson(w, http.StatusBadRequest, models.APIError{
// 			Error: err,
// 		})
// 	}
// 	priceResponse := models.PriceResponse{
// 		Price:  price,
// 		Ticker: ticker,
// 	}
// 	return utils.WriteJson(w, http.StatusAccepted, priceResponse)
// }
