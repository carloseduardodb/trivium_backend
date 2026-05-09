package controller

import (
	"net/http"

	"trivium/internal/domain/usecase"
	"trivium/internal/presentation/format"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

type PortfolioController struct {
	portfolioUseCase *usecase.PortfolioUseCase
	authMiddleware   *middleware.Auth
}

func NewPortfolioController(portfolioUseCase *usecase.PortfolioUseCase, authMiddleware *middleware.Auth) *PortfolioController {
	return &PortfolioController{
		portfolioUseCase: portfolioUseCase,
		authMiddleware:   authMiddleware,
	}
}

func (c *PortfolioController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	protected := router.SubRouter("/portfolio")
	protected.Use(c.authMiddleware.AuthMiddleware())

	protected.HandleFunc("", c.GetPortfolio, http.MethodGet)
}

func (c *PortfolioController) GetPortfolio(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserID(r.Context())
	if userId == 0 {
		format.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	portfolio, err := c.portfolioUseCase.GetPortfolio(userId)
	if err != nil {
		format.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	format.WriteSuccessResponse(w, http.StatusOK, portfolio)
}
