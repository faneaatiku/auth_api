package service

import (
	"fmt"
	"github.com/faneaatiku/auth_api/app/dto"
	"github.com/faneaatiku/auth_api/app/entity"
	"github.com/labstack/echo/v4"
)

type UserStorage interface {
}

type AuthService struct {
	logger      echo.Logger
	userStorage UserStorage
}

func NewAuthService(logger echo.Logger, userStorage UserStorage) (*AuthService, error) {
	if logger == nil || userStorage == nil {
		return nil, fmt.Errorf("invalid dependencies provided to AuthService constructor")
	}

	return &AuthService{
		logger:      logger,
		userStorage: userStorage,
	}, nil
}

func (as AuthService) RegisterUser(request dto.RegisterRequest) (entity.User, error) {
	return entity.User{Email: "treaba mea"}, nil
}
