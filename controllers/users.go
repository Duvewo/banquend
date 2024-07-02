package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/models"
	"github.com/labstack/echo/v4"
)

type UsersController struct {
	handler.Handler
}

func (c UsersController) REGISTER(group *echo.Group) {
	group.GET("/:query", c.ByQuery)
}

func (c UsersController) ByQuery(ctx echo.Context) error {
	target := models.UserModel{}

	if err := ctx.Bind(&target); err != nil {
		return fmt.Errorf("users/byquery: to bind: %w", err)
	}

	user, err := c.Users.Search(context.Background(), target)

	if err != nil {
		return fmt.Errorf("users/byquery: to search: %w", err)
	}

	return ctx.JSON(http.StatusOK, user)
}
