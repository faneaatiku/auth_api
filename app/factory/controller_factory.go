package factory

import (
	"github.com/faneaatiku/auth_api/app/controller"
	"github.com/labstack/echo/v4"
)

func GetAuthController(logger echo.Logger) (*controller.AuthController, error) {
	return controller.NewAuthController(logger)
}
