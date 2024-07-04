package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/models"
	"github.com/labstack/echo/v4"
)

type CurrenciesController struct {
	handler.Handler
}

func (c CurrenciesController) REGISTER(g *echo.Group) {
	g.GET("/:code", c.ByCode)
}

func (c CurrenciesController) ByCode(ctx echo.Context) error {
	cur, err := c.Currencies.ByCode(context.Background(), models.CurrencyModel{Code: ctx.Param("code")})

	if err != nil {
		return fmt.Errorf("currencies/byCode: to byCode(), %w", err)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"ok":   true,
		"data": cur,
	})

}
