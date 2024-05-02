package main

import (
	"embed"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ChoregateUI is the embedded UI files.
//
//go:embed public/*
var choregateUI embed.FS

func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the static files from the embedded filesystem
	http.FileServer(http.FS(choregateUI)).ServeHTTP(w, r)
}

func main() {
	// Configure the logger level and format
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.RFC3339Nano

	var signalHandler = make(chan os.Signal, 1)
	var shutdown = make(chan bool, 1)

	log.Info().Msg("starting choregate ui")

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/public", http.StatusSeeOther)
	})
	r.Get("/public/*", staticHandler)

	go func() {
		s := <-signalHandler
		log.Info().Msgf("received signal: %v", s)
		shutdown <- true
	}()

	srv := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("choregate ui error")
			signalHandler <- syscall.SIGTERM
		}
	}()

	// Wait for a signal
	<-shutdown
	log.Info().Msg("shutting down choregate ui")
}
