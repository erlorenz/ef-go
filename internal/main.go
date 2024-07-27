package main

import (
	"errors"
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

	mux := http.NewServeMux()
	r := ef.New(mux)

	r.Get("/something", SomeHandler)
	r.Get("/somethingbad", SomeErrorHandler)

	r.Mux.HandleFunc("GET /somethingfunc", ef.HF(SomeHandler))
	r.Mux.HandleFunc("GET /somethingbadfunc", ef.HF(SomeErrorHandler))

	app.logger.Info("Server listening...", "port", ":4000")
	http.ListenAndServe(":4000", r.Mux)
}

func SomeHandler(w http.ResponseWriter, r *http.Request) error {
	return ef.JSON(w, 200, "Some JSON message")
}

func SomeErrorHandler(w http.ResponseWriter, r *http.Request) error {
	err := errors.New("some error")
	return ef.ErrorJSON(err, 402, "This the is human error message.")
}
