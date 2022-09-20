// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"
	"time"

	"github.com/ardanlabs/bets/app/services/engine/handlers/v1/brunogrp"
	"github.com/ardanlabs/bets/app/services/engine/handlers/v1/gamegrp"
	"github.com/ardanlabs/bets/business/core/bank"
	"github.com/ardanlabs/bets/business/web/auth"
	"github.com/ardanlabs/bets/business/web/v1/mid"
	"github.com/ardanlabs/bets/foundation/events"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log            *zap.SugaredLogger
	Auth           *auth.Auth
	Converter      *currency.Converter
	Bank           *bank.Bank
	Evts           *events.Events
	AnteUSD        float64
	BankTimeout    time.Duration
	ConnectTimeout time.Duration
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	// Register group endpoints.
	ggh := gamegrp.Handlers{
		Converter:      cfg.Converter,
		Bank:           cfg.Bank,
		Log:            cfg.Log,
		Evts:           cfg.Evts,
		WS:             websocket.Upgrader{},
		Auth:           cfg.Auth,
		BankTimeout:    cfg.BankTimeout,
		ConnectTimeout: cfg.ConnectTimeout,
	}

	app.Handle(http.MethodPost, version, "/game/connect", ggh.Connect)
	app.Handle(http.MethodGet, version, "/game/events", ggh.Events)
	app.Handle(http.MethodGet, version, "/game/config", ggh.Configuration)
	app.Handle(http.MethodGet, version, "/game/usd2wei/:usd", ggh.USD2Wei)
	app.Handle(http.MethodGet, version, "/game/test", ggh.Test, mid.Authenticate(cfg.Log, cfg.Auth))

	var bgh brunogrp.Handlers

	app.Handle(http.MethodGet, version, "/game/bruno/bets/:page/:rows", bgh.Query)
	app.Handle(http.MethodGet, version, "/game/bruno/bet/:id", bgh.QueryByID)
	app.Handle(http.MethodPost, version, "/game/bruno/bet/:id", bgh.Create)
	app.Handle(http.MethodPost, version, "/game/bruno/signBet", bgh.SignBet)
	app.Handle(http.MethodPost, version, "/game/bruno/setWinner", bgh.AcceptMod)

}
