package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	handler.Handler
}

func (c AuthController) REGISTER(group *echo.Group) {

}
