package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type AccountsController struct {
	handler.Handler
}

func (c AccountsController) REGISTER(group *echo.Group) {

}
