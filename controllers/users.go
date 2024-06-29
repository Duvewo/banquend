package controllers

import (
	"github.com/Duvewo/banquend/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersController struct {
	*pgxpool.Pool
	handler.Handler
	Path string
}

func (c UsersController) REGISTER(h handler.Handler) {
	c.Router.Group(c.Path)
	c.Router.GET("/:query", nil)
	c.Router.POST("/create", nil)
}
