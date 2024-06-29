package controllers

import (
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
	group.POST("/create", c.Create)
}

func (c UsersController) ByQuery(ctx echo.Context) error {
	return nil
}

func (c UsersController) Create(ctx echo.Context) error {
	user := models.UserModel{}
	if err := c.Router.Binder.Bind(&user, ctx); err != nil {
		c.Logger.Errorf("controllers/users: to bind %w", err)

		return err

	}

	c.Logger.Infoln(user)

	// user.Validate()

	// err := c.Users.Create(user)
	return ctx.JSON(http.StatusCreated, echo.Map{"ok": true})
}
