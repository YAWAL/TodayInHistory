package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/YAWAL/TodayInHistory/app"
	"github.com/YAWAL/TodayInHistory/logging"
)

func main() {

	file := flag.String("config", "config.json", "config file path")
	flag.Parse()
	srv, err := app.LoadApp(*file)
	if err != nil {
		logging.Log.Errorf("error during starting app: %v", err)
	}

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go app.GracefullShutdown(srv, quit, done)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.Log.Errorf("Could not listen on %s: %v\n", os.Args[0], err)
	}
	logging.Log.Info("Server stopped")
	<-done
}
