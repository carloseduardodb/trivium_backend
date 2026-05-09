package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/dto"
	"trivium/internal/presentation/format"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/mux"
)

type ProfitTakeController struct {
	profitTakeUseCase *usecase.ProfitTakeUseCase
	authMiddleware    *middleware.Auth
}

func NewProfitTakeController(profitTakeUseCase *usecase.ProfitTakeUseCase, authMiddleware *middleware.Auth) *ProfitTakeController {
	return &ProfitTakeController{
		profitTakeUseCase: profitTakeUseCase,
		authMiddleware:    authMiddleware,
	}
}

func (c *ProfitTakeController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/profit-takes")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("", c.Create, http.MethodPost)
	protected.HandleFunc("/position/{positionId}", c.FindByPosition, http.MethodGet)
	protected.HandleFunc("/{id}", c.Delete, http.MethodDelete)
}

func (c *ProfitTakeController) Create(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var input dto.CreateProfitTake
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	profitTake, err := c.profitTakeUseCase.Create(userId, &input)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusCreated, profitTake)
}

func (c *ProfitTakeController) FindByPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	positionId, err := strconv.ParseInt(vars["positionId"], 10, 64)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid position ID")
		return
	}

	profitTakes, err := c.profitTakeUseCase.FindByPositionId(positionId)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, profitTakes)
}

func (c *ProfitTakeController) Delete(w http.ResponseWriter, r *http.Request) {
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

	if err := c.profitTakeUseCase.Delete(id, userId); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, map[string]string{"message": "deleted successfully"})
}
