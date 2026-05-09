package cmd

import (
	"os"

	"trivium/internal/presentation/controller"
	"trivium/internal/presentation/middleware"
	presentation_repositorier "trivium/internal/presentation/repositorier"
	"trivium/internal/presentation/route"

	"github.com/google/wire"
)

type ServerStarter interface {
	Start() error
}

type AppServer struct {
	AuthController           *controller.AuthController
	StatusController         *controller.StatusController
	CryptoCurrencyController *controller.CryptoCurrencyController
	PositionController       *controller.PositionController
	ProfitTakeController     *controller.ProfitTakeController
	PortfolioController      *controller.PortfolioController
	UserController           *controller.UserController
	CryptoHistoryController  *controller.CryptoHistoryController
	PriceAlertController     *controller.PriceAlertController
	WsCryptoController       *controller.WsCryptoController
	Router                   presentation_repositorier.HttpRepositorier
}

func NewAppServer(
	auth *controller.AuthController,
	status *controller.StatusController,
	crypto *controller.CryptoCurrencyController,
	position *controller.PositionController,
	profitTake *controller.ProfitTakeController,
	portfolio *controller.PortfolioController,
	user *controller.UserController,
	cryptoHistory *controller.CryptoHistoryController,
	priceAlert *controller.PriceAlertController,
	wsCrypto *controller.WsCryptoController,
	router presentation_repositorier.HttpRepositorier,
) *AppServer {
	return &AppServer{
		AuthController:           auth,
		StatusController:         status,
		CryptoCurrencyController: crypto,
		PositionController:       position,
		ProfitTakeController:     profitTake,
		PortfolioController:      portfolio,
		UserController:           user,
		CryptoHistoryController:  cryptoHistory,
		PriceAlertController:     priceAlert,
		WsCryptoController:       wsCrypto,
		Router:                   router,
	}
}

func (a *AppServer) Start() error {
	port := os.Getenv("PORT")
	return route.NewRoutes(
		port,
		a.Router,
		a.AuthController,
		a.StatusController,
		a.CryptoCurrencyController,
		a.PositionController,
		a.ProfitTakeController,
		a.PortfolioController,
		a.UserController,
		a.CryptoHistoryController,
		a.PriceAlertController,
		a.WsCryptoController,
	)
}

var PresentationModule = wire.NewSet(
	controller.NewAuthController,
	controller.NewStatusController,
	controller.NewCryptoCurrencyController,
	controller.NewPositionController,
	controller.NewProfitTakeController,
	controller.NewPortfolioController,
	controller.NewUserController,
	controller.NewCryptoHistoryController,
	controller.NewPriceAlertController,
	controller.NewWsCryptoController,
	middleware.NewAuth,
	NewAppServer,
	wire.Bind(new(ServerStarter), new(*AppServer)),
)
