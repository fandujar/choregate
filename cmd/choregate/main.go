package main

import (
	"embed"
	"net/http"

	"github.com/fandujar/golaze"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// go:embed ui/dist
var uiFS embed.FS

func main() {
	uiRouter := UiRouter(uiFS)

	ui := golaze.NewWebApp(
		&golaze.WebAppConfig{
			Port:   "8080",
			Router: uiRouter,
		},
	)

	app := golaze.NewApp(
		&golaze.AppConfig{
			Name:     "choregate",
			LogLevel: zerolog.DebugLevel,
			WebApp:   ui,
		},
	)

	if err := app.Run(); err != nil {
		log.Error().Err(err).Msg("error running app")
	}
}

// UiRouter return a router with the routes for the UI
func UiRouter(fs embed.FS) *chi.Mux {
	router := golaze.NewRouter()
	router.Handle("/*", http.FileServer(http.FS(fs)))

	return router
}
