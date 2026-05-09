package controller

import (
	"net/http"

	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/format"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

type UserController struct {
	userRepo       repositorier.UserRepositorier
	authMiddleware *middleware.Auth
}

func NewUserController(userRepo repositorier.UserRepositorier, authMiddleware *middleware.Auth) *UserController {
	return &UserController{
		userRepo:       userRepo,
		authMiddleware: authMiddleware,
	}
}

func (c *UserController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/users")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("/me", c.Me, http.MethodGet)
}

func (c *UserController) Me(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := c.userRepo.FindById(userId)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, user)
}
