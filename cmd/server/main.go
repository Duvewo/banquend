package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"

	"github.com/Duvewo/banquend/controllers"
	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/storage"
	"github.com/Duvewo/banquend/storage/postgres"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	FLAG_SRV_ADDR  = flag.String("srv-addr", os.Getenv("SRV_ADDR"), "Server address to listen")
	FLAG_IS_SECURE = flag.Bool("secure", true, "Run server in secure mode")
)

// TODO: generate random
var JWT_KEY = []byte("jwtkeyhere")

func main() {
	flag.Parse()
	e := echo.New()
	log, err := zap.NewDevelopment()

	if err != nil {
		//TODO: handle this error better
		os.Exit(1)
	}

	sugar := log.Sugar()
	sugar.Debugln("logging is ready")

	e.Debug = !*FLAG_IS_SECURE

	pg, err := postgres.Open(context.Background(), "")

	if err != nil {
		os.Exit(1)
	}

	//TODO: handle this
	if err := pg.Ping(context.Background()); err != nil {
		sugar.Infoln("postgres does not respond")
	}

	h := handler.Handler{
		Router: e,

		Users:    &storage.Users{Pool: pg},
		Payments: &storage.Payments{Pool: pg},

		Logger: sugar,
	}

	api := h.Router.Group("/api")
	/* AUTH MANAGEMENT */
	auth := api.Group("/auth")

	/* USER MANAGEMENT */
	user := api.Group("/users")
	//friends := user.Group("/friends")
	accounts := user.Group("/accounts")

	/* PAYMENT MANAGEMENT */
	payments := api.Group("/payments")
	cards := payments.Group("/cards")
	// crypto := payments.Group("/crypto")

	/* SUPPORT MANAGEMENT */
	support := api.Group("/support")

	controllers.Register(&controllers.AuthController{Handler: h}, auth)
	sugar.Debugln(h.Router.Routes())

	controllers.Register(&controllers.UsersController{Handler: h}, user)
	controllers.AccountsController{Handler: h}.REGISTER(accounts)

	controllers.PaymentsController{Handler: h}.REGISTER(payments)
	controllers.CardsController{Handler: h}.REGISTER(cards)
	//crypto

	controllers.SupportController{Handler: h}.REGISTER(support)

	sugar.Debugln("starting a server...")
	if err := h.Router.Start(*FLAG_SRV_ADDR); !errors.Is(err, http.ErrServerClosed) {
		sugar.Errorf("to start a server: ", err)
	}

}
