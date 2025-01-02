package controller

import (
	"net/http"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

type StatusController struct{}

func NewStatusController() *StatusController {
	return &StatusController{}
}

func (s *StatusController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	router.HandleFunc("/check-status", s.Status, http.MethodGet)
}

func (s *StatusController) Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
