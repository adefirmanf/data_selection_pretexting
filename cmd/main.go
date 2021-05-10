package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/adefirmanf/go-api-template/config"
	"github.com/adefirmanf/go-api-template/config/env"
	"github.com/adefirmanf/go-api-template/internal/httpserver"
	"github.com/adefirmanf/go-api-template/internal/metricsserver"
)

func main() {
	envConfig := env.New()
	config.Init(envConfig)

	cfg := config.Load()

	// app := app.New()
	// Todo : Handling signal exit
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func(c context.Context) {
		h := httpserver.Httpserver{
			Metrics: httpserver.Metrics{},
		}
		httpAppInit := httpserver.NewHTTPServer(h)
		fmt.Printf("Service http-app-listener started [:%v] \n", cfg.AppPort())
		httpAppInit.ListenAndServe(cfg.AppPort())
	}(ctx)

	go func(c context.Context) {
		httpMetricsInit := metricsserver.NewHTTPServer()
		fmt.Printf("Service http-metrics-listener started [:%v] \n", cfg.MetricsPort())
		httpMetricsInit.ListenAndServe(cfg.MetricsPort())
	}(ctx)

	<-quit
	cancel()
}
