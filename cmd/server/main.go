package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"

	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/service"
	"github.com/Duvewo/banquend/storage"
	"github.com/Duvewo/banquend/storage/postgres"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	FLAG_SRV_ADDR = flag.String("srv-addr", os.Getenv("SRV_ADDR"), "Server address to listen")
)

func main() {
	e := echo.New()
	log, err := zap.NewProduction()

	if err != nil {
		//TODO: handle this error better
		os.Exit(1)
	}

	sugar := log.Sugar()
	sugar.Debugln("logging is ready")

	pg, err := postgres.Open(context.Background(), "")

	if err != nil {
		os.Exit(1)
	}

	h := handler.Handler{
		Router:   e,
		Users:    &storage.Users{Pool: pg},
		Payments: &storage.Payments{Pool: pg},
	}

	api := h.Router.Group("/api")
	/* AUTH MANAGEMENT */
	api.Group("/auth")

	/* USER MANAGEMENT */
	user := api.Group("/user")
	service.REGISTER()
	//friends := user.Group("/friends")

	/* PAYMENT MANAGEMENT */
	payments := api.Group("/payments")
	payments.Group("/crypto")

	/* SUPPORT MANAGEMENT */
	api.Group("/support")

	sugar.Debugln("starting a server...")
	if err := h.Router.Start(*FLAG_SRV_ADDR); !errors.Is(err, http.ErrServerClosed) {
		sugar.Errorf("to start a server: ", err)
	}

}
