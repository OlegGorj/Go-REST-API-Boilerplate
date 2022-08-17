package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/auth"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/config"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/errors"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/global"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/health"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/internal/project"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/accesslog"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/dbcontext"
	"github.com/OlegGorj/Go-REST-API-Boilerplate/pkg/log"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

func main() {

	// connect to the database
	db, err := dbx.MustOpen("postgres", global.AppConfig.DSN)
	if err != nil {
		global.Logger.Error(err)
		os.Exit(-1)
	}
	db.QueryLogFunc = logDBQuery(global.Logger)
	db.ExecLogFunc = logDBExec(global.Logger)
	defer func() {
		if err := db.Close(); err != nil {
			global.Logger.Error(err)
		}
	}()

	// build HTTP server
	address := fmt.Sprintf(":%v", global.AppConfig.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(global.Logger, dbcontext.New(db), global.AppConfig),
	}

	// start the HTTP server with graceful shutdown
	go routing.GracefulShutdown(hs, 10*time.Second, global.Logger.Infof)
	global.Logger.Infof("server %v is running at %v", global.Version, address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.Logger.Error(err)
		os.Exit(-1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, db *dbcontext.DB, cfg *config.Config) http.Handler {
	router := routing.New()
	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)

	health.RegisterHandlers(router, global.Version)

	rg := router.Group(fmt.Sprintf("/%s", global.APIGroup))

	authHandler := auth.Handler(cfg.JWTSigningKey)

	project.RegisterHandlers(rg.Group(""),
		project.NewService(project.NewRepository(db, logger), logger),
		authHandler, logger,
	)

	auth.RegisterHandlers(rg.Group(""),
		auth.NewService(cfg.JWTSigningKey, cfg.JWTExpiration, logger),
		logger,
	)

	return router
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(logger log.Logger) dbx.QueryLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB query error: %v", err)
		}
	}
}

// logDBExec returns a logging function that can be used to log SQL executions.
func logDBExec(logger log.Logger) dbx.ExecLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB execution successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB execution error: %v", err)
		}
	}
}
