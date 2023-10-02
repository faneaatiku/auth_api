package controller

import (
	"fmt"
	"github.com/faneaatiku/auth_api/app/dto"
	"github.com/faneaatiku/iac"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	logger echo.Logger
}

func NewAuthController(logger echo.Logger) (*AuthController, error) {
	if logger == nil {
		return nil, fmt.Errorf("invalid dependencies provided to AuthController")
	}

	return &AuthController{logger: logger}, nil
}

func (c *AuthController) HandleRegistration(ctx echo.Context) error {
	var req dto.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, iac.BuildResponse(iac.WithGenericMsg("invalid request")))
	}

	return ctx.JSON(http.StatusOK, iac.BuildResponse(iac.WithGenericMsg("Hello, World!")))
}
