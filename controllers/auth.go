package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/internal/jwt"
	"github.com/Duvewo/banquend/models"
	jwtgo "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	handler.Handler
}

func (c *AuthController) REGISTER(group *echo.Group) {
	group.POST("/register", c.Register)
	group.POST("/login", c.Login)
	group.POST("/restore", c.Restore)
	c.Logger.Debugf("registered authcontroller, %v\n", c.Router.Routes()[0])
}

func (c *AuthController) Register(ctx echo.Context) error {
	userForm := models.UserModel{}

	if err := ctx.Bind(&userForm); err != nil {
		return fmt.Errorf("to bind: %w", err)
	}

	//TODO: VALIDATE DATA

	if err := c.Users.Create(context.Background(), userForm); err != nil {
		return fmt.Errorf("to create: %w", err)
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"ok": true})
}

func (c *AuthController) Login(ctx echo.Context) error {
	var authModel models.AuthModel
	if err := ctx.Bind(&authModel); err != nil {
		return fmt.Errorf("to bind: %w", err)
	}

	c.Users.Search(
		context.Background(),
		models.UserModel{
			PhoneNumber: authModel.PhoneNumber,
			Email:       authModel.Email,
			Password:    authModel.Password,
		})

	claims := jwt.AuthClaims{
		AuthModel: models.AuthModel{
			ID: 1,
		},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS384, claims)
	signed, err := token.SignedString([]byte("key"))

	if err != nil {
		return fmt.Errorf("to sign: %w", err)
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{"ok": true, "token": signed})
}

func (c *AuthController) Restore(ctx echo.Context) error {
	return nil
}
