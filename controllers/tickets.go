package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type TicketsController struct {
	handler.Handler
}

func (c TicketsController) REGISTER(group *echo.Group) {

}
