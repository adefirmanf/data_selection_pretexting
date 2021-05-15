package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/adefirmanf/data_selection_pretexting/app"
	"github.com/adefirmanf/data_selection_pretexting/config"
	"github.com/adefirmanf/data_selection_pretexting/config/env"
	"github.com/adefirmanf/data_selection_pretexting/internal/httpserver"
	"github.com/adefirmanf/data_selection_pretexting/internal/jobsserver"
	"github.com/adefirmanf/data_selection_pretexting/internal/metricsserver"
	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"

	_ "github.com/lib/pq"
)

func main() {
	initApp := app.New()
	envConfig := env.New()
	config.Init(envConfig)

	cfg := config.Load()

	config := jobsserver.NewConfig(1, "Seconds")
	jobserver := jobsserver.NewJobServer(config, initApp.Storage, initApp.Queue)
	urls := jobsserver.NewQueryURLS()
	urls.Add(scrappertweets.NewQueryURL("JeniusConnect", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("bankmandiri", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("linkaja", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("kontakBRI", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("ovo_id", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("BNI", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("HaloBCA", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("danawallet", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))
	urls.Add(scrappertweets.NewQueryURL("gopayindonesia", "Whatsapp|Mohon Maaf|LiveChat", "is:reply"))

	scrapperTweets := jobsserver.NewScrapperTweets(jobserver, initApp.ScrapperTweets, urls)
	scrapperUsers := jobsserver.NewScrapperUsers(jobserver, initApp.ScrapperUsers)
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
	go func(c context.Context) {
		fmt.Printf("Job user scrapper service started \n")
		jobstartUser(initApp, scrapperUsers)
		defer fmt.Println("Done")
	}(ctx)
	go func(c context.Context) {
		jobStartTweets(scrapperTweets, 50)
		defer fmt.Println("Done")
	}(ctx)

	<-quit
	cancel()
}

func jobstartUser(app *app.App, scrapperUsers *jobsserver.ScrapperUsers) {
	for {
		len := app.Queue.Size()
		if len > 0 {
			UserID := app.Queue.PullFront()
			err := scrapperUsers.Scrape(UserID)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

func jobStartTweets(scrapperTweets *jobsserver.ScrapperTweets, maxBatch int) {
	err := scrapperTweets.Scrape(maxBatch)
	if err != nil {
		log.Panic(err)
	}
}
