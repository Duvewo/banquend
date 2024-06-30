package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type UsersController struct {
	handler.Handler
}

func (c UsersController) REGISTER(group *echo.Group) {
	group.GET("/:query", c.ByQuery)
}

func (c UsersController) ByQuery(ctx echo.Context) error {
	return nil
}
