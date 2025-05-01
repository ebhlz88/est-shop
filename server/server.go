package server

import (
	"github.com/ebhlz88/est-shop/common"
)

type JSONApiServer struct {
	ListenAddr string
	Store      common.Store
}
type APIError struct {
	Error error
}

type APISuccessMessage struct {
	Message string
}

func NewJSONApiServer(listenAddr string, store common.Store) *JSONApiServer {
	return &JSONApiServer{
		ListenAddr: listenAddr,
		Store:      store,
	}
}

func (s *JSONApiServer) GetStore() common.Store {
	return s.Store
}
