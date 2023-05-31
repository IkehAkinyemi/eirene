package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IkehAkinyemi/myblog/internal/models"
	"github.com/IkehAkinyemi/myblog/internal/util"
	"github.com/rs/zerolog"
)

type Server struct {
	configs util.Configs
	logger zerolog.Logger
}

func NewServer(configs util.Configs,
	store models.Store,
) *Server {
	return &Server{
		configs: configs,
	}
}

func (s *Server) Start() error {
	srv := &http.Server{
		Addr: s.configs.ServerAddress,
		Handler: s.setupRoutes(),
		ErrorLog:     log.New(s.logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownErr := make(chan error)

	// Background job to listen for any shutdown signal
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sc := <-quit

		s.logger.Info().
		Str("signal", sc.String()).
		Msg("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownErr <- err
		}

		s.logger.Info().
			Str("addr", srv.Addr).
			Msg("completing background tasks")

		shutdownErr <- nil
	}()

	s.logger.Info().
		Str("environment", s.configs.Env).
		Str("addr", srv.Addr).
		Msg("starting server")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return err
	}

	s.logger.Info().
		Str("addr", srv.Addr).
		Msg("server stopped")

	return nil
}
