package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pcpratheesh/go-urlshortner/config"
	"github.com/pcpratheesh/go-urlshortner/controllers"
	"github.com/pcpratheesh/go-urlshortner/docs"
	"github.com/pcpratheesh/go-urlshortner/middleware"
	"github.com/sirupsen/logrus"

	_ "github.com/pcpratheesh/go-urlshortner/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

var (
	APP_NAME = "URLSRT"
)

// @title Go URL Shortner App
// @version 2.0
// @description This is a golang url shortner application.

// @host localhost
// @BasePath /
// @schemes http
func main() {

	app := fiber.New()

	// load the configurations
	cfg, err := config.LoadConfigurations(APP_NAME)
	if err != nil {
		logrus.Fatal(err)
	}

	// initialize the controller
	controller, err := controllers.NewController(controllers.Config{
		BaseURL: cfg.BaseURL,
		Store:   cfg.Store,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	{
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Println("Server gracefully shutting down...")
			_ = app.Shutdown()
		}()
	}

	{

		swaggerHost := strings.Replace(cfg.BaseURL, "https://", "", -1)
		swaggerHost = strings.Replace(swaggerHost, "http://", "", -1)
		// SET swagger info
		docs.SwaggerInfo.Title = "Go URL Shortner Docs"
		docs.SwaggerInfo.Description = "Go URL Shortner API Docs"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = swaggerHost
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		// endpoint handler declarations
		app.Get("/", controller.Home)
		app.Get("/health", controller.Health)
		app.Post("/encode", middleware.ApiKeyMiddleware(cfg.XApiKey), controller.EncodeURL)
		app.Get("/short/:url", controller.DecodeURL)

		app.Get("/swagger/*", fiberSwagger.WrapHandler)

	}

	// server run
	if err := app.Listen(fmt.Sprintf(":%v", cfg.Port)); err != nil {
		log.Panic(err)
	}

}
