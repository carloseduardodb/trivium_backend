// route/route.go
package route

import (
	"fmt"
	"trivium/internal/presentation/controller"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

func NewRoutes(
	router presentation_repositorier.HttpRepositorier,
	authController *controller.AuthController,
	statusController *controller.StatusController,
) error {
	authController.SetupRoutes(router)
	statusController.SetupRoutes(router)

	fmt.Println("Servidor iniciado em http://127.0.0.1:3000")
	return router.Start("127.0.0.1:3000")
}
