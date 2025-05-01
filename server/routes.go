package server

import (
	"net/http"

	"github.com/ebhlz88/est-shop/handler"
	"github.com/ebhlz88/est-shop/utils"
)

func Run(s *JSONApiServer) {
	// http.HandleFunc("/pricefetch", utils.MakeHTTPHandlerFunc(handler.HandlePriceFetch))
	http.HandleFunc("/", utils.MakeHTTPHandlerFunc(handler.HandleProduct(s.Store)))
	http.HandleFunc("/product", utils.MakeHTTPHandlerFunc(handler.HandleProductById(s.Store)))
	http.ListenAndServe(s.ListenAddr, nil)
}
