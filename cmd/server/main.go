package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/Duvewo/banquend/controllers"
	"github.com/Duvewo/banquend/handler"
	"github.com/Duvewo/banquend/internal/generator"
	"github.com/Duvewo/banquend/internal/jwt"
	"github.com/Duvewo/banquend/storage"
	"github.com/Duvewo/banquend/storage/postgres"
	jwtgo "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	FLAG_SRV_ADDR   = flag.String("srv-addr", os.Getenv("SRV_ADDR"), "Server address to listen")
	FLAG_CACHE_ADDR = flag.String("cache-addr", os.Getenv("CACHE_ADDR"), "Cache address to listen")
	FLAG_DB_ADDR    = flag.String("db-addr", os.Getenv("DB_ADDR"), "Database address to listen")
	FLAG_IS_SECURE  = flag.Bool("secure", os.Getenv("MODE_SECURE") == "secure", "Run server in secure mode")
)

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

	// cache, err := redis.Open(*FLAG_CACHE_ADDR)

	// if err != nil {
	// 	sugaredLogger.Fatalf("open redis: %v", err)
	// }

	// if err := cache.Ping(context.Background()).Err(); err != nil {
	// 	sugaredLogger.Fatalf("to ping cache: %v", err)
	// }

	h := handler.Handler{
		Router: router,
		//Cache:    cache,
		Logger: sugaredLogger,
		JWTKey: generator.MustGenerate(
			int(generator.MustGenerateNumber(big.NewInt(3333)).Int64() + 100),
		),
		Users:      &storage.Users{Pool: db},
		Accounts:   &storage.Accounts{Pool: db},
		Currencies: &storage.Currencies{Pool: db},
		Payments:   &storage.Payments{Pool: db},
	}

	h.Logger.Debugf("key is %v\n", h.JWTKey)

	api := h.Router.Group("/api")

	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Bearer JWT
			bearer := c.Request().Header.Get("Authorization")
			if bearer != "" {
				token := strings.Split(bearer, " ")[1]
				h.Logger.Debugln(token)
				parsedToken, err := jwtgo.ParseWithClaims(token, &jwt.AuthClaims{}, func(t *jwtgo.Token) (interface{}, error) {
					return h.JWTKey, nil
				}, jwtgo.WithValidMethods([]string{jwtgo.SigningMethodHS384.Name}))
				if err != nil {
					return fmt.Errorf("mware: %w", err)
				}

				h.Logger.Debugf("Parsed token: %s\n", parsedToken.Raw)
				c.Set("user", parsedToken)
			}
			return next(c)
		}
	})

	/* AUTH MANAGEMENT */
	auth := api.Group("/auth")

	/* USER MANAGEMENT */
	user := api.Group("/users")
	//friends := user.Group("/friends")
	accounts := user.Group("/accounts")
	currencies := accounts.Group("/currencies")

	/* PAYMENT MANAGEMENT */
	payments := api.Group("/payments")
	cards := payments.Group("/cards")
	// crypto := payments.Group("/crypto")

	/* SUPPORT MANAGEMENT */
	support := api.Group("/support")

	controllers.Register(&controllers.AuthController{Handler: h}, auth)

	controllers.Register(&controllers.UsersController{Handler: h}, user)
	controllers.Register(&controllers.AccountsController{Handler: h}, accounts)
	controllers.Register(&controllers.CurrenciesController{Handler: h}, currencies)

	controllers.Register(&controllers.PaymentsController{Handler: h}, payments)
	controllers.Register(&controllers.CardsController{Handler: h}, cards)

	controllers.Register(&controllers.SupportController{Handler: h}, support)

	h.Logger.Debugln("controllers registered")

	h.Logger.Debugln("starting a server...")
	if err := h.Router.Start(*FLAG_SRV_ADDR); !errors.Is(err, http.ErrServerClosed) {
		h.Logger.Errorf("to start a server: ", err)
	}

}
