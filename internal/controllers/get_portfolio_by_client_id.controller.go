package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/camilo-cpp/golang-api-echo/internal/dtos"
	"github.com/camilo-cpp/golang-api-echo/internal/services"
)

func GetPortfolioByClientId(c echo.Context) error {
	clientId := c.Param("clientId")

	client := &services.GetPortfolioByClientIdClient{}

	portfolioData, err := client.GetPortfolioByClientIdService(clientId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, portfolioData)
}

func GetPortfolioItemsByClientId(c echo.Context) error {
	portfolioId := c.Param("portfolioId")
	pageSize := c.QueryParam("pageSize")
	currentPage := c.QueryParam("currentPage")

	params := &dtos.ParamsGetPortfolioItemsByClientId{
		PortfolioId: portfolioId,
		PageSize:    pageSize,
		CurrentPage: currentPage,
	}

	client := &services.GetPortfolioByClientIdClient{}

	itemsData, err := client.GetPortfolioItemsByClientIdService(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, itemsData)

}
