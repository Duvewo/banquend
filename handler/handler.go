package handler

import (
	"github.com/Duvewo/banquend/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	Router   *echo.Echo
	Logger   *zap.SugaredLogger
	Users    storage.UsersRepository
	Payments storage.PaymentsRepository
}
