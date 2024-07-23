package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/erlorenz/ef-framework/ef"
)

type App struct {
	logger *slog.Logger
}

func main() {

	app := App{logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}

	router := ef.New(app.logger)

	router.Get("/something", app.HandleSomething)

	app.logger.Info("Server listening...", "port", ":4000")
	http.ListenAndServe(":4000", router)
}
