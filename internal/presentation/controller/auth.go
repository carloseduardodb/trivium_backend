package controller

import (
	"fmt"
	"net/http"
	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/dto"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

type AuthController struct {
	authUsecase    *usecase.AuthUseCase
	authMiddleware middleware.Auth
}

func NewAuthController(authUsecase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		authUsecase: authUsecase,
	}
}

func (a *AuthController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	router.HandleFunc("/auth", middleware.JsonMiddleware(a.Auth, &dto.Auth{}), http.MethodPost)

	protected := router.SubRouter("/auth")
	protected.Use(a.authMiddleware.AuthMiddleware())
}

func (c *AuthController) Auth(input interface{}) (interface{}, error) {
	req, ok := input.(*dto.Auth)
	if !ok {
		return nil, fmt.Errorf("invalid input format")
	}

	return c.authUsecase.Auth(req)
}
