package server

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/book-library/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Set(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(logger.LoggingMiddleware)

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	r.Use(cors.Handler(corsOptions))
}

func SetConfig(dirpath string, filename string) {
	filePath := filepath.Join(dirpath, filename)
	fileExist := isFileExist(filePath)

	if fileExist {
		viper.AddConfigPath(dirpath)
		viper.SetConfigFile(filePath)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal().Err(err).Msgf("error reading config file: %+v", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	SecretConfig()
}

// isFileExist check if the file exist on the given file path
func isFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
