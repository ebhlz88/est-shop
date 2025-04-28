package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type JSONApiServer struct {
	listenAddr string
	svc        PriceFetcher
	store      PostgresStore
}
type APIError struct {
	Error error
}

type APISuccessMessage struct {
	Message string
}

func NewJSONApiServer(listenAddr string, svc PriceFetcher, store PostgresStore) *JSONApiServer {
	return &JSONApiServer{
		listenAddr: listenAddr,
		svc:        svc,
		store:      store,
	}
}

func (s *JSONApiServer) Run() {
	http.HandleFunc("/pricefetch", MakeHTTPHandlerFunc(s.handlePriceFetch))
	http.HandleFunc("/", MakeHTTPHandlerFunc(s.handleProduct))
	http.HandleFunc("/product", MakeHTTPHandlerFunc(s.handleProductById))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONApiServer) handleProduct(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.HandleGetProducts(w, r)
	}
	if r.Method == "POST" {
		return s.HandlePostProducts(w, r)
	}
	return fmt.Errorf("%s Method not allowd", r.Method)
}

func (s *JSONApiServer) handleProductById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	if r.Method == "GET" {
		return s.HandleGetProductById(w, r, id)
	}

	if r.Method == "DELETE" {
		return nil
	}
	if r.Method == http.MethodPatch {
		return s.handlePatchProduct(w, r, id)
	}
	return fmt.Errorf("%s Method not allowd", r.Method)
}

func (s *JSONApiServer) handlePatchProduct(w http.ResponseWriter, r *http.Request, id int) error {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return WriteJson(w, http.StatusInternalServerError, APIError{
			Error: err,
		})
	}
	if err := s.store.ModifyProduct(id, product.ProductName, product.ProductBuyPrice, product.ProductSellPrice); err != nil {
		return WriteJson(w, http.StatusInternalServerError, APIError{
			Error: err,
		})
	}
	return WriteJson(w, http.StatusOK, product)
}
func (s *JSONApiServer) HandleGetProductById(w http.ResponseWriter, r *http.Request, id int) error {
	product, err := s.store.GetProductById(id)
	if err != nil {
		return WriteJson(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}
	return WriteJson(w, http.StatusOK, product)
}

func (s *JSONApiServer) HandleGetProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := s.store.GetAllProducts()
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, APIError{
			Error: err,
		})
	}
	return WriteJson(w, http.StatusAccepted, products)
}

func (s *JSONApiServer) HandlePostProducts(w http.ResponseWriter, r *http.Request) error {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return WriteJson(w, http.StatusBadRequest, APIError{
			Error: err,
		})
	}
	err := s.store.AddProduct(product.ProductName, product.ProductBuyPrice, product.ProductSellPrice)
	if err != nil {
		return WriteJson(w, http.StatusInternalServerError, APIError{
			Error: err,
		})
	}
	return WriteJson(w, http.StatusCreated, map[string]string{
		"ok": "product successfully added",
	})
}

func (s *JSONApiServer) HandleDeleteProduct(w http.ResponseWriter, r *http.Request, id int) error {
	err := s.store.DeleteProduct(id)
	if err != nil {
		return WriteJson(w, http.StatusInternalServerError, APIError{
			Error: err,
		})
	}
	return WriteJson(w, http.StatusNoContent, APISuccessMessage{
		Message: "Delete Successful",
	})
}

func (s *JSONApiServer) HandlePatchProduct(w http.ResponseWriter, r *http.Request, id int) error {

	return WriteJson(w, http.StatusOK, APISuccessMessage{
		Message: "Succesfully Updated",
	})
}

func (s *JSONApiServer) handlePriceFetch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")
	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, APIError{
			Error: err,
		})
	}
	priceResponse := PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return WriteJson(w, http.StatusAccepted, priceResponse)
}

type ApiFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func MakeHTTPHandlerFunc(apiFunc ApiFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", rand.Intn(1000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			WriteJson(w, http.StatusAccepted, APIError{
				Error: err,
			})
		}
	}
}

func WriteJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}

func GetId(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("productId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Errorf("%s is not found")
	}
	return id, err
}
