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

type PositionController struct {
	positionUseCase *usecase.PositionUseCase
	authMiddleware  *middleware.Auth
}

func NewPositionController(positionUseCase *usecase.PositionUseCase, authMiddleware *middleware.Auth) *PositionController {
	return &PositionController{
		positionUseCase: positionUseCase,
		authMiddleware:  authMiddleware,
	}
}

func (c *PositionController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/positions")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("", c.FindByUser, http.MethodGet)
	protected.HandleFunc("/{id}", c.FindById, http.MethodGet)
	protected.HandleFunc("", c.Create, http.MethodPost)
	protected.HandleFunc("/{id}/close", c.Close, http.MethodPatch)
	protected.HandleFunc("/{id}", c.Delete, http.MethodDelete)
}

func (c *PositionController) Create(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var input dto.CreatePosition
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	position, err := c.positionUseCase.Create(userId, &input)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusCreated, position)
}

func (c *PositionController) FindByUser(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	positions, err := c.positionUseCase.FindByUserId(userId)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, positions)
}

func (c *PositionController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	position, err := c.positionUseCase.FindById(id)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, position)
}

func (c *PositionController) Close(w http.ResponseWriter, r *http.Request) {
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

	position, err := c.positionUseCase.Close(id, userId)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, position)
}

func (c *PositionController) Delete(w http.ResponseWriter, r *http.Request) {
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

	if err := c.positionUseCase.Delete(id, userId); err != nil {
		format.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, map[string]string{"message": "deleted successfully"})
}
