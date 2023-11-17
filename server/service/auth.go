package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

type RequestAuthenticator struct {
	authorizedApps map[string]interface{}
}

func NewRequestAuthenticator() (*RequestAuthenticator, error) {
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

	return &RequestAuthenticator{
		authorizedApps: authorized,
	}, nil
}

func (as *RequestAuthenticator) ValidateAppKey(key string, c echo.Context) (bool, error) {
	_, ok := as.authorizedApps[key]
	if !ok {
		return false, fmt.Errorf("unauthorized application")
	}

	return true, nil
}
