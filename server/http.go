package server

import (
	"github.com/faneaatiku/auth_api/connector"
	"github.com/faneaatiku/auth_api/server/factory"
	"github.com/faneaatiku/auth_api/server/service"
	"github.com/faneaatiku/iac"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Start() {
	// Echo instance
	e := echo.New()
	authService, err := service.NewRequestAuthenticator()
	if err != nil {
		e.Logger.Fatal("could not start server: %s", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//generates a unique id for each request
	e.Use(middleware.RequestID())
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:app-key",
		ErrorHandler: func(err error, c echo.Context) error {
			e.Logger.Infof("error on authentication: %s", err)
			return c.JSON(http.StatusUnauthorized, iac.BuildResponse(iac.WithGenericMsg("unauthorized")))
		},
		Validator: authService.ValidateAppKey,
	}))

	//add custom error handler in order to return the response formatted as we want
	e.HTTPErrorHandler = getGenericErrorHandler(e)

	connections, err := connector.NewConnections()
	if err != nil {
		e.Logger.Fatalf("could not start server: %s", err)
	}

	ctrlFactory, err := factory.NewControllerFactory(connections, e.Logger)
	if err != nil {
		e.Logger.Fatalf("could not start server: %s", err)
	}

	authController, err := ctrlFactory.GetAuthController()
	if err != nil {
		e.Logger.Fatalf("could not start server: %s", err)
	}

	// Routes
	e.POST("/register", authController.HandleRegistration)

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}

func getGenericErrorHandler(e *echo.Echo) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		e.Logger.Warnf("generic error handler called: %s", err)
		httpError, ok := err.(*echo.HTTPError)
		if !ok {
			_ = c.JSON(http.StatusInternalServerError, iac.BuildResponse(iac.WithGenericMsg("an unknown error occurred")))
			return
		}

		_ = c.JSON(httpError.Code, iac.BuildResponse(iac.WithGenericMsg(httpError.Message.(string))))
	}
}
