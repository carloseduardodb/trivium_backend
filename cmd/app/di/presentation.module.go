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
	AuthController   *controller.AuthController
	StatusController *controller.StatusController
	Router           presentation_repositorier.HttpRepositorier
}

func NewAppServer(auth *controller.AuthController, status *controller.StatusController, router presentation_repositorier.HttpRepositorier) *AppServer {
	return &AppServer{
		AuthController:   auth,
		StatusController: status,
		Router:           router,
	}
}

func (a *AppServer) Start() error {
	port := os.Getenv("PORT")
	return route.NewRoutes(port, a.Router, a.AuthController, a.StatusController)
}

var PresentationModule = wire.NewSet(
	controller.NewAuthController,
	controller.NewStatusController,
	middleware.NewAuth,
	NewAppServer,
	wire.Bind(new(ServerStarter), new(*AppServer)),
)
