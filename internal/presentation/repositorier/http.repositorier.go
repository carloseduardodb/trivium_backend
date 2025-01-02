package presentation_repositorier

import "net/http"

type HttpRepositorier interface {
	HandleFunc(path string, handler http.HandlerFunc, method string)
	SubRouter(pathPrefix string) HttpRepositorier
	Use(middleware func(http.Handler) http.Handler)
	Start(address string) error
}
