package controller

import (
	"net/http"

	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/format"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/mux"
)

type CryptoHistoryController struct {
	cryptoHistoryRepo repositorier.CryptoHistoryRepository
	volumeRepo        repositorier.VolumeRepository
	authMiddleware    *middleware.Auth
}

func NewCryptoHistoryController(
	cryptoHistoryRepo repositorier.CryptoHistoryRepository,
	volumeRepo repositorier.VolumeRepository,
	authMiddleware *middleware.Auth,
) *CryptoHistoryController {
	return &CryptoHistoryController{
		cryptoHistoryRepo: cryptoHistoryRepo,
		volumeRepo:        volumeRepo,
		authMiddleware:    authMiddleware,
	}
}

func (c *CryptoHistoryController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/crypto")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("/history", c.FindAllHistory, http.MethodGet)
	protected.HandleFunc("/history/{symbol}", c.FindHistoryBySymbol, http.MethodGet)
	protected.HandleFunc("/volume/{symbol}", c.FindVolumeBySymbol, http.MethodGet)
}

func (c *CryptoHistoryController) FindAllHistory(w http.ResponseWriter, r *http.Request) {
	history, err := c.cryptoHistoryRepo.FindAll()
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, history)
}

func (c *CryptoHistoryController) FindHistoryBySymbol(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]

	history, err := c.cryptoHistoryRepo.FindBySymbol(symbol)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, history)
}

func (c *CryptoHistoryController) FindVolumeBySymbol(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]

	volumes, err := c.volumeRepo.FindBySymbol(symbol)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, volumes)
}
