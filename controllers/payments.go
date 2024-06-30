package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type PaymentsController struct {
	handler.Handler
}

func (c PaymentsController) REGISTER(group *echo.Group) {

}
