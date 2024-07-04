package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type CardsController struct {
	handler.Handler
}

func (c CardsController) REGISTER(group *echo.Group) {
	c.Router.GET("/:account_id", c.ByAccountID)
	c.Router.GET("/:card_id/:action", c.ActionByCardID)
	c.Router.DELETE("/:card_id", c.DeleteByID)
}

func (c CardsController) ByAccountID(ctx echo.Context) error {
	return nil
}

func (c CardsController) ActionByCardID(ctx echo.Context) error {
	return nil
}

func (c CardsController) DeleteByID(ctx echo.Context) error {
	return nil
}
