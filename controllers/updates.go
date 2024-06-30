package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type UpdatesController struct {
	handler.Handler
}

func (c UpdatesController) REGISTER(group *echo.Group) {

}
