package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ebhlz88/est-shop/common"
	"github.com/ebhlz88/est-shop/models"
)

func MakeHTTPHandlerFunc(apiFunc common.ApiFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", rand.Intn(1000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			WriteJson(w, http.StatusAccepted, models.APIError{
				Error: err.Error(),
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
