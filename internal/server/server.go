package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/camilo-cpp/golang-api-echo/internal/routes"
)

func Start() error {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	routes.GetPortfolioByClientId(e)
	routes.GetPortfolioItemsByClientId(e)

	port := os.Getenv("PORT")

	if err := e.Start(":" + port); err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}

	return nil
}
