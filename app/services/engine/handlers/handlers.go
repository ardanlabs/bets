// Package handlers manages the different versions of the API.
package handlers

import (
	"context"
	"expvar"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/ardanlabs/bets/app/services/engine/handlers/debug/checkgrp"
	v1 "github.com/ardanlabs/bets/app/services/engine/handlers/v1"
	"github.com/ardanlabs/bets/business/core/book"
	"github.com/ardanlabs/bets/business/web/auth"
	"github.com/ardanlabs/bets/business/web/v1/mid"
	"github.com/ardanlabs/bets/foundation/events"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown       chan os.Signal
	Log            *zap.SugaredLogger
	Auth           *auth.Auth
	Converter      *currency.Converter
	Book           *book.Book
	Evts           *events.Events
	AnteUSD        float64
	BankTimeout    time.Duration
	ConnectTimeout time.Duration
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	// Construct the web.App which holds all routes as well as common Middleware.
	var app *web.App

	// Do we need CORS?
	if opts.corsOrigin != "" {
		app = web.NewApp(
			cfg.Shutdown,
			mid.Logger(cfg.Log),
			mid.Errors(cfg.Log),
			mid.Metrics(),
			mid.Cors(opts.corsOrigin),
			mid.Panics(),
		)

		// Accept CORS 'OPTIONS' preflight requests if config has been provided.
		// Don't forget to apply the CORS middleware to the routes that need it.
		// Example Config: `conf:"default:https://MY_DOMAIN.COM"`
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			return nil
		}
		app.Handle(http.MethodOptions, "", "/*", h, mid.Cors(opts.corsOrigin))
	}

	if app == nil {
		app = web.NewApp(
			cfg.Shutdown,
			mid.Logger(cfg.Log),
			mid.Errors(cfg.Log),
			mid.Metrics(),
			mid.Panics(),
		)
	}

	// Load the v1 routes.
	v1.Routes(app, v1.Config{
		Log:            cfg.Log,
		Auth:           cfg.Auth,
		Converter:      cfg.Converter,
		Book:           cfg.Book,
		Evts:           cfg.Evts,
		AnteUSD:        cfg.AnteUSD,
		BankTimeout:    cfg.BankTimeout,
		ConnectTimeout: cfg.ConnectTimeout,
	})

	return app
}

// DebugStandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}

// DebugMux registers all the debug standard library routes and then custom
// debug application routes for the service. This bypassing the use of the
// DefaultServerMux. Using the DefaultServerMux would be a security risk since
// a dependency could inject a handler into our service without us knowing it.
func DebugMux(build string, log *zap.SugaredLogger, db *sqlx.DB) http.Handler {
	mux := DebugStandardLibraryMux()

	// Register debug check endpoints.
	cgh := checkgrp.Handlers{
		Build: build,
		Log:   log,
		DB:    db,
	}
	mux.HandleFunc("/debug/readiness", cgh.Readiness)
	mux.HandleFunc("/debug/liveness", cgh.Liveness)

	return mux
}
