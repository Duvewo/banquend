package controllers

import (
	"context"
	"fmt"
	"net/http"

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
	c.Router.DELETE("/:id/close", c.CloseByID)
	c.Router.PATCH("/:id/restrict", c.RestrictByID)

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

	if err := c.Accounts.Create(context.Background(), account); err != nil {
		return fmt.Errorf("accountsController/Create(): to create, %w", err)
	}
	return ctx.JSON(http.StatusOK, echo.Map{"ok": true})
}

func (c AccountsController) CloseByID(ctx echo.Context) error {
	return nil
}

func (c AccountsController) RestrictByID(ctx echo.Context) error {
	return nil
}
