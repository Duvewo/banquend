package handler

import (
	"github.com/Duvewo/banquend/storage"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Router   *echo.Echo
	Users    storage.UsersRepository
	Payments storage.PaymentsRepository
}
