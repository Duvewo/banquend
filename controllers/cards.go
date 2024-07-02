package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type CardsController struct {
	handler.Handler
}

func (c CardsController) REGISTER(group *echo.Group) {
	c.Router.GET("/:number", c.ByNumber)
}

func (c CardsController) ByNumber(ctx echo.Context) error {
	return nil
}
