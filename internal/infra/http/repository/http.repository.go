package repository

import (
	"net/http"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/mux"
)

type HttpRepository struct {
	router *mux.Router
}

func NewHttpRepository() presentation_repositorier.HttpRepositorier {
	return &HttpRepository{router: mux.NewRouter()}
}

func (m *HttpRepository) HandleFunc(path string, handler http.HandlerFunc, method string) {
	m.router.HandleFunc(path, handler).Methods(method)
}

func (m *HttpRepository) SubRouter(pathPrefix string) presentation_repositorier.HttpRepositorier {
	return &HttpRepository{router: m.router.PathPrefix(pathPrefix).Subrouter()}
}

func (m *HttpRepository) Use(middleware func(http.Handler) http.Handler) {
	m.router.Use(middleware)
}

func (m *HttpRepository) Start(address string) error {
	return http.ListenAndServe(address, m.router)
}
