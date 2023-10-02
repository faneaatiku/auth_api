package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

type AuthService struct {
	logger echo.Logger

	//the key on which we store the AuthorizedApp is the PublicKey
	authorizedApps map[string]interface{}
}

func NewAuthService(logger echo.Logger) (*AuthService, error) {
	if logger == nil {
		return nil, fmt.Errorf("invalid dependencies provided to AuthService constructor")
	}

	authorized := make(map[string]interface{})
	envKeys, found := os.LookupEnv("AUTHORIZED_KEYS")
	if found {
		split := strings.Split(envKeys, ";")
		for _, key := range split {
			key = strings.TrimSpace(key)
			if key == "" {
				continue
			}
			authorized[key] = nil
		}
	}

	return &AuthService{
		logger:         logger,
		authorizedApps: authorized,
	}, nil
}

func (as *AuthService) ValidateAppKey(key string, c echo.Context) (bool, error) {
	_, ok := as.authorizedApps[key]
	if !ok {
		return false, fmt.Errorf("unauthorized application")
	}

	return true, nil
}
