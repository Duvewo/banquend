package controllers

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	REGISTER(*echo.Group)
}

func Register(c Controller, g *echo.Group) {
	c.REGISTER(g)
}
