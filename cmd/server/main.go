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
	"github.com/Duvewo/banquend/storage/redis"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	FLAG_SRV_ADDR   = flag.String("srv-addr", os.Getenv("SRV_ADDR"), "Server address to listen")
	FLAG_CACHE_ADDR = flag.String("cache-addr", os.Getenv("CACHE_ADDR"), "Cache address to listen")
	FLAG_DB_ADDR    = flag.String("db-addr", os.Getenv("DB_ADDR"), "Database address to listen")
	FLAG_IS_SECURE  = flag.Bool("secure", true, "Run server in secure mode")
)

// TODO: generate random
var JWT_KEY = []byte("jwtkeyhere")

func main() {
	flag.Parse()
	router := echo.New()
	log, err := zap.NewDevelopment()

	if err != nil {
		//TODO: handle this error better
		os.Exit(1)
	}

	sugaredLogger := log.Sugar()
	sugaredLogger.Debugln("logging is ready")

	router.Debug = !*FLAG_IS_SECURE

	db, err := postgres.Open(context.Background(), *FLAG_DB_ADDR)

	if err != nil {
		sugaredLogger.Fatalf("to open: %v", err)
	}

	//TODO: handle this
	if err := db.Ping(context.Background()); err != nil {
		sugaredLogger.Fatalf("to ping: %v", err)
	}

	cache, err := redis.Open(*FLAG_CACHE_ADDR)

	if err != nil {
		sugaredLogger.Fatalf("open redis: %v", err)
	}

	if err := cache.Ping(context.Background()).Err(); err != nil {
		sugaredLogger.Fatalf("to ping cache: %v", err)
	}

	h := handler.Handler{
		Router:   router,
		Cache:    cache,
		Logger:   sugaredLogger,
		Users:    &storage.Users{Pool: db},
		Payments: &storage.Payments{Pool: db},
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

	controllers.Register(&controllers.UsersController{Handler: h}, user)
	controllers.Register(&controllers.AccountsController{Handler: h}, accounts)

	controllers.Register(&controllers.PaymentsController{Handler: h}, payments)
	controllers.Register(&controllers.CardsController{Handler: h}, cards)

	controllers.Register(&controllers.SupportController{Handler: h}, support)

	h.Logger.Debugln("controllers registered")

	h.Logger.Debugln("starting a server...")
	if err := h.Router.Start(*FLAG_SRV_ADDR); !errors.Is(err, http.ErrServerClosed) {
		h.Logger.Errorf("to start a server: ", err)
	}

}
