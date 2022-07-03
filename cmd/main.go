package main

import (
	"Casino/config"
	"Casino/internal/server"
	"Casino/pkg/logger"
	"Casino/pkg/storage/postgres"
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Starting server")
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	log.Println("Config loaded")

	appLogger := logger.NewApiLogger(cfg)
	err = appLogger.InitLogger()
	if err != nil {
		log.Fatalf("Cannot init logger: %v", err.Error())
	}
	appLogger.Infof("Logger successfully started with - Level: %s, InFile: %t (filePath: %s), InTG: %t (chatID: %d)",
		cfg.Logger.Level,
		cfg.Logger.InFile,
		cfg.Logger.FilePath,
		cfg.Logger.InTG,
		cfg.Logger.ChatID,
	)

	appLogger.Infof("Server config - Host: %s, Port: %s",
		cfg.Server.Host,
		cfg.Server.Port,
	)

	if cfg.Server.ShowUnknownErrorsInResponse {
		appLogger.Warnf("ShowUnknownErrorsInResponse: %t", cfg.Server.ShowUnknownErrorsInResponse)
	}

	pgDB, err := postgres.InitPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
		return
	}
	appLogger.Infof("Postgres connected, Status: %#v", pgDB.Stats())

	defer func(pgDB *sqlx.DB) {
		err = pgDB.Close()
		if err != nil {
			appLogger.Infof(err.Error())
		} else {
			appLogger.Info("PostgreSQL closed properly")
		}
	}(pgDB)

	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.OpenTelemetry.URL)))
	if err != nil {
		appLogger.Fatalf("Cannot create Jaeger exporter - %s", err.Error())
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.OpenTelemetry.ServiceName),
		)),
	)

	otel.SetTracerProvider(tp)
	defer func() {
		err = tp.Shutdown(context.Background())
		if err != nil {
			appLogger.Error(err)
		}
	}()

	s := server.NewServer(cfg, pgDB, appLogger)
	if err = s.Run(); err != nil {
		appLogger.Fatalf("Cannot start server: %v", err)
	}
}
