package handler

import (
	"github.com/Duvewo/banquend/storage"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Handler struct {
	Router     *echo.Echo
	Cache      *redis.Client
	Logger     *zap.SugaredLogger
	JWTKey     []byte
	Users      storage.UsersRepository
	Accounts   storage.AccountsRepository
	Cards      storage.CardsRepository
	Auth       storage.AuthRepository
	Currencies storage.CurrenciesRepository
	Payments   storage.PaymentsRepository
}
