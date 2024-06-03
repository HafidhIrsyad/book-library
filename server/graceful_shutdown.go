package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func startServerWithGracefulShutdown(r *chi.Mux) {
	addr := fmt.Sprintf(":%d", AppPort)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Create server context
	serverCtx, cancelServerCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		graceful_timeout := 10

		shutDownCtx, cancel := context.WithTimeout(serverCtx, time.Duration(graceful_timeout)*time.Second)
		defer cancel()

		go func() {
			<-shutDownCtx.Done()
			if shutDownCtx.Err() == context.DeadlineExceeded {
				log.Fatal().Err(shutDownCtx.Err()).Msg("graceful shutdown timed out, forcing exit..")
			}
		}()

		err := server.Shutdown(shutDownCtx)
		if err != nil {
			log.Fatal().Err(err).Msgf("error on shutting down gracefully: %v", err)
		}

		cancelServerCtx()
	}()

	log.Info().Msgf("starting book-library in port: %d", AppPort)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msgf("error on starting up server: %v", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	log.Info().Msg("server is shut down!")
}
