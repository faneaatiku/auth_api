package controller

import (
	"fmt"
	"github.com/faneaatiku/auth_api/app/dto"
	"github.com/faneaatiku/auth_api/app/entity"
	"github.com/faneaatiku/iac"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthService interface {
	RegisterUser(request dto.RegisterRequest) (entity.User, error)
}

type AuthController struct {
	logger echo.Logger

	authService AuthService
}

func NewAuthController(logger echo.Logger, authService AuthService) (*AuthController, error) {
	if logger == nil || authService == nil {
		return nil, fmt.Errorf("invalid dependencies provided to AuthController")
	}

	return &AuthController{logger: logger, authService: authService}, nil
}

func (c *AuthController) HandleRegistration(ctx echo.Context) error {
	var req dto.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, iac.BuildResponse(iac.WithGenericMsg("invalid request")))
	}

	err := req.Validate()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, iac.BuildResponse(iac.WithGenericMsg(err.Error())))
	}

	_, err = c.authService.RegisterUser(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, iac.BuildResponse(iac.WithGenericMsg(err.Error())))
	}

	return ctx.JSON(http.StatusOK, iac.BuildResponse())
}
