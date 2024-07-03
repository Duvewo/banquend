package controllers

import (
	"context"
	"fmt"

	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/models"
	"github.com/labstack/echo/v4"
)

type AccountsController struct {
	handler.Handler
}

func (c AccountsController) REGISTER(group *echo.Group) {
	c.Router.GET("/:id", c.ByID)
	c.Router.GET("/:owner_id", c.ByOwnerID)

	c.Router.POST("/create", c.Create)
	c.Router.DELETE("/:id/close", c.Close)

}

func (c AccountsController) ByID(ctx echo.Context) error {
	return nil
}

func (c AccountsController) ByOwnerID(ctx echo.Context) error {
	return nil
}

func (c AccountsController) Create(ctx echo.Context) error {
	var account models.AccountModel

	if err := ctx.Bind(&account); err != nil {
		return fmt.Errorf("to bind: %w", err)
	}

	return c.Accounts.Create(context.Background(), account)
}

func (c AccountsController) Close(ctx echo.Context) error {
	return nil
}
