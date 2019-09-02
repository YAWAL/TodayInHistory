package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/YAWAL/TodayInHistory/config"
	"github.com/YAWAL/TodayInHistory/database"
	"github.com/YAWAL/TodayInHistory/logging"
	"github.com/YAWAL/TodayInHistory/model"
	"github.com/YAWAL/TodayInHistory/repository"
	"github.com/YAWAL/TodayInHistory/router"
)

// LoadApp performs reading of  config file, initializes connection to database,
// performs initial table migrations, initializes Postgres repository, initializes routers
// and creates server
func LoadApp(file string) (srv *http.Server, err error) {
	// read, config
	conf, err := config.ReadConfig(file)
	if err != nil {
		logging.Log.Errorf("cannot load config: %v", err.Error())
		return nil, err
	}
	// establish connections to DB
	db, err := database.PGconn(conf.Database)
	if err != nil {
		logging.Log.Errorf("cannot connect to DB: %s", err.Error())
		return nil, err
	}
	// migrate tables
	db.AutoMigrate(&model.Link{}, &model.Event{}, &model.Birth{}, &model.Death{})
	logging.Log.Infof("Auto migration")

	//init repos
	repo := repository.NewPostgresHistoryRepository(db)
	logging.Log.Infof("Initializing Postgres repository")

	// init routers
	r := router.InitRouter(repo)
	logging.Log.Infof("Initializing routers")

	srv = &http.Server{
		Handler:      r,
		Addr:         conf.Host,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	logging.Log.Infof("Application is running on %s ", conf.Host)

	return srv, nil
}

// GracefullShutdown performs grecefull shutdown of server and loggs to stdout
// info about operation success
func GracefullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logging.Log.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logging.Log.Errorf("can not gracefully shutdown the server: %v", err)
	}
	close(done)
}
