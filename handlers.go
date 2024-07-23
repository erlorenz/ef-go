package main

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/erlorenz/ef-framework/ef"
)

func (app App) HandleSomething(ctx ef.Context, w http.ResponseWriter, r *http.Request) (ef.Output, error) {
	body := map[string]string{
		"status":  "in progress",
		"started": "2024-05-05 10:00:00Z",
	}

	int := rand.Intn(4)
	delay := rand.Intn(300)

	ctx.Logger().Info("Randoms", "int", int, "delay", delay)

	time.Sleep(time.Millisecond * time.Duration(delay))

	if int == 1 {
		return ef.ErrorJSON(errors.New("int is 1"), 500, "error: int is 1")
	}

	if int == 2 {
		return ef.ErrorHTML(errors.New("int is 2"), 200, "error: int is 2")
	}

	if int == 3 {
		return ef.HTML(201, `<h1>Hello world!</h1>`)
	}

	return ef.JSON(201, body)
}
