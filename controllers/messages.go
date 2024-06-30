package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type MessagesController struct {
	handler.Handler
}

func (c MessagesController) REGISTER(group *echo.Group) {

}
