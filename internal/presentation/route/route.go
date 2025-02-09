// route/route.go
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
) error {
	authController.SetupRoutes(router)
	statusController.SetupRoutes(router)

	if port == "" {
		port = "3000"
	}

	fmt.Println("Servidor iniciado na porta:", port)
	return router.Start("127.0.0.1:" + port)
}
