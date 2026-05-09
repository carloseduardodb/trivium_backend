package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/format"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/mux"
)

type PriceAlertController struct {
	alertUseCase   *usecase.PriceAlertUseCase
	authMiddleware *middleware.Auth
}

func NewPriceAlertController(alertUseCase *usecase.PriceAlertUseCase, authMiddleware *middleware.Auth) *PriceAlertController {
	return &PriceAlertController{
		alertUseCase:   alertUseCase,
		authMiddleware: authMiddleware,
	}
}

func (c *PriceAlertController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/alerts")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("", c.Create, http.MethodPost)
	protected.HandleFunc("", c.FindByUser, http.MethodGet)
	protected.HandleFunc("/{id}", c.Delete, http.MethodDelete)
}

func (c *PriceAlertController) Create(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var input usecase.CreateAlertInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	alert, err := c.alertUseCase.Create(userId, &input)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusCreated, alert)
}

func (c *PriceAlertController) FindByUser(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	alerts, err := c.alertUseCase.FindByUserId(userId)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, alerts)
}

func (c *PriceAlertController) Delete(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := c.alertUseCase.Delete(id, userId); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, map[string]string{"message": "deleted successfully"})
}
