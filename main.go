package main

import (
	"github.com/hugovantighem/utrade/api"
	"github.com/hugovantighem/utrade/infra/db"
	"github.com/hugovantighem/utrade/infra/handlers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	srv := echo.New()

	repo := db.NewInMemoryRepository()

	loggingApi, err := handlers.NewApi(repo)
	if err != nil {
		logrus.Errorf("cannot build api: %v", err)
		return
	}

	api.RegisterHandlers(srv, loggingApi)

	srv.Logger.Fatal(srv.Start(":8080")) // TODO graceful shutdown

}
