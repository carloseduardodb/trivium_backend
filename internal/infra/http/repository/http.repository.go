package repository

import (
	"net/http"

	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/mux"
)

type HttpRepository struct {
	router *mux.Router
}

func NewHttpRepository() presentation_repositorier.HttpRepositorier {
	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)
	r.Use(middleware.LoggingMiddleware)
	return &HttpRepository{router: r}
}

func (m *HttpRepository) HandleFunc(path string, handler http.HandlerFunc, method string) {
	m.router.HandleFunc(path, handler).Methods(method, http.MethodOptions)
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
