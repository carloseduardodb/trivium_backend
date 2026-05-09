package route

import (
	"fmt"

	"trivium/internal/presentation/controller"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

func NewRoutes(
	port string,
	router presentation_repositorier.HttpRepositorier,
	authController *controller.AuthController,
	statusController *controller.StatusController,
	cryptoCurrencyController *controller.CryptoCurrencyController,
	positionController *controller.PositionController,
	profitTakeController *controller.ProfitTakeController,
	portfolioController *controller.PortfolioController,
	userController *controller.UserController,
	cryptoHistoryController *controller.CryptoHistoryController,
	priceAlertController *controller.PriceAlertController,
	wsCryptoController *controller.WsCryptoController,
) error {
	authController.SetupRoutes(router)
	statusController.SetupRoutes(router)
	cryptoCurrencyController.SetupRoutes(router)
	positionController.SetupRoutes(router)
	profitTakeController.SetupRoutes(router)
	portfolioController.SetupRoutes(router)
	userController.SetupRoutes(router)
	cryptoHistoryController.SetupRoutes(router)
	priceAlertController.SetupRoutes(router)
	wsCryptoController.SetupRoutes(router)

	if port == "" {
		port = "3000"
	}

	fmt.Println("Server started on port:", port)
	return router.Start("0.0.0.0:" + port)
}
