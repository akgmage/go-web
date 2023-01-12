package main

import (
	"net/http"

	"github.com/akgmage/go-web/pkg/config"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	return mux
}