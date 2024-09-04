package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/camilo-cpp/golang-api-echo/internal/controllers"

)

func GetPortfolioByClientId(e *echo.Echo) {
	e.GET("/portfolio/client/:clientId", controllers.GetPortfolioByClientId)
}

func GetPortfolioItemsByClientId(e *echo.Echo) {
	e.GET("/portfolio/items/client/:portfolioId", controllers.GetPortfolioItemsByClientId)
}
