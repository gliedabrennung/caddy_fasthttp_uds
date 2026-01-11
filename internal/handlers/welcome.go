package handlers

import (
	"encoding/json"
	"log/slog"
	"time"

	"github.com/valyala/fasthttp"
)

type WelcomeHandler struct {
	Welcome     string `json:"welcome"`
	FullDate    string `json:"full_date"`
	TimeOnly    string `json:"time_only"`
	CurrentYear int    `json:"current_year"`
}

func WelcomePage(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	data := WelcomeHandler{
		Welcome:     "Go APP работает",
		FullDate:    time.Now().Format("02.01.2006"),
		TimeOnly:    time.Now().Format("15:04"),
		CurrentYear: time.Now().Year(),
	}

	err := json.NewEncoder(ctx.Response.BodyWriter()).Encode(data)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return
	}
}
