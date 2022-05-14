package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mashbens/my-movie-list/api"
	"github.com/mashbens/my-movie-list/app/modules"
	"github.com/mashbens/my-movie-list/config"
	"github.com/mashbens/my-movie-list/util"

	_ "github.com/mashbens/my-movie-list/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Movie API
// @version 1.0
// @description Brikut API yang digunakan untuk memanage movie data
func main() {
	config := config.GetConfig()
	dbCon := util.NewConnectionDatabase(config)

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	handleSwag := echoSwagger.WrapHandler

	e.GET("/swagger/*", handleSwag)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	api.RegisterRoutes(e, &controllers)

	go func() {
		addres := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(addres); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	defer dbCon.CloseConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}
