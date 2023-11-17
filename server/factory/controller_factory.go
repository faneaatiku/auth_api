package factory

import (
	"fmt"
	"github.com/faneaatiku/auth_api/app/controller"
	"github.com/faneaatiku/auth_api/app/repository"
	"github.com/faneaatiku/auth_api/app/service"
	"github.com/faneaatiku/auth_api/connector"
	"github.com/labstack/echo/v4"
)

type ControllerFactory struct {
	conn   *connector.Connections
	logger echo.Logger
}

func NewControllerFactory(conn *connector.Connections, logger echo.Logger) (*ControllerFactory, error) {
	if conn == nil || logger == nil {
		return nil, fmt.Errorf("could not instantiate controller factory: invalid dependencies")
	}

	return &ControllerFactory{
		conn:   conn,
		logger: logger,
	}, nil
}

func (c *ControllerFactory) GetAuthController() (*controller.AuthController, error) {
	userRepo, err := repository.NewUserRepository(c.conn.Mysql, c.logger)
	if err != nil {
		return nil, err
	}

	authService, err := service.NewAuthService(c.logger, userRepo)

	return controller.NewAuthController(c.logger, authService)
}
