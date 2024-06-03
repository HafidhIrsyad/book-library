package logger

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogConfig struct {
	Logger  zerolog.Logger
	Req     *http.Request
	Ctx     context.Context
	Request interface{}
	Err     error
	Message string
}

func LogInfo(log LogConfig) (l zerolog.Logger) {
	log.Logger.Info().
		Ctx(log.Ctx).
		Str("method", log.Req.Method).
		Str("path", log.Req.URL.Path).
		Str("host", log.Req.Host).
		Str("user-agent", log.Req.UserAgent()).
		Msgf("Request Received: %v", log.Request)

	return log.Logger
}

func LogError(log LogConfig) (l zerolog.Logger) {
	log.Logger.Error().Err(log.Err).
		Ctx(log.Ctx).
		Str("method", log.Req.Method).
		Str("path", log.Req.URL.Path).
		Str("host", log.Req.Host).
		Str("user-agent", log.Req.UserAgent()).
		Msgf("message: %s - request: %v", log.Message, log.Request)

	return log.Logger
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
