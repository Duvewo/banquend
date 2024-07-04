package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/labstack/echo/v4"
)

type UsersController struct {
	handler.Handler
}

func (c UsersController) REGISTER(group *echo.Group) {
	// group.GET("/:public_id", c.ByPublicID)
	// group.GET("/:email", c.ByEmail)
	// group.GET("/:phone_number", c.ByPhone)
}
