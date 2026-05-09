package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/dto"
	"trivium/internal/presentation/format"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/mux"
)

type CryptoCurrencyController struct {
	cryptoUseCase  *usecase.CryptoCurrencyUseCase
	authMiddleware *middleware.Auth
}

func NewCryptoCurrencyController(cryptoUseCase *usecase.CryptoCurrencyUseCase, authMiddleware *middleware.Auth) *CryptoCurrencyController {
	return &CryptoCurrencyController{
		cryptoUseCase:  cryptoUseCase,
		authMiddleware: authMiddleware,
	}
}

func (c *CryptoCurrencyController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/cryptocurrencies")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("", c.FindAll, http.MethodGet)
	protected.HandleFunc("/{id}", c.FindById, http.MethodGet)
	protected.HandleFunc("", c.Create, http.MethodPost)
	protected.HandleFunc("/{id}", c.Update, http.MethodPut)
	protected.HandleFunc("/{id}", c.Delete, http.MethodDelete)
}

func (c *CryptoCurrencyController) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCryptoCurrency
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	crypto, err := c.cryptoUseCase.Create(&input)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusCreated, crypto)
}

func (c *CryptoCurrencyController) FindAll(w http.ResponseWriter, r *http.Request) {
	cryptos, err := c.cryptoUseCase.FindAll()
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, cryptos)
}

func (c *CryptoCurrencyController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	crypto, err := c.cryptoUseCase.FindById(id)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, crypto)
}

func (c *CryptoCurrencyController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input dto.UpdateCryptoCurrency
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}
	input.ID = id

	crypto, err := c.cryptoUseCase.Update(&input)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, crypto)
}

func (c *CryptoCurrencyController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := c.cryptoUseCase.Delete(id); err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error deleting cryptocurrency: %v", err))
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, map[string]string{"message": "deleted successfully"})
}
