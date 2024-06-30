package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type TransactionsController struct {
	handler.Handler
}

func (c TransactionsController) REGISTER(group *echo.Group) {

}
