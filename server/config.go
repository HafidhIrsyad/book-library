package server

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type dsnConfig struct {
	host     string
	user     string
	password string
	db       string
	port     string
}

var (
	DriverName              string
	DbHost                  string
	DbUser                  string
	DbPassword              string
	DbName                  string
	DbPort                  string
	AppPort                 int
	DbMaxOpenConnection     int
	DbMaxIdleConnection     int
	DbConnectionMaxLifeTime time.Duration
)

func SecretConfig() {
	DriverName = viper.GetString("DRIVER_NAME")
	DbHost = viper.GetString("DB_HOST")
	DbUser = viper.GetString("DB_USER")
	DbPassword = viper.GetString("DB_PASSWORD")
	DbName = viper.GetString("DB_NAME")
	DbPort = viper.GetString("DB_PORT")

	AppPort = viper.GetInt("APP_PORT")
	DbMaxOpenConnection = viper.GetInt("DB_MAX_OPEN_CONNECTION")
	DbMaxIdleConnection = viper.GetInt("DB_MAX_IDLE_CONNECTION")

	viper.SetDefault("DB_CONNECTION_MAX_LIFE_TIME", time.Second*time.Duration(300))
	DbConnectionMaxLifeTime = time.Second * time.Duration(viper.GetInt("DB_CONNECTION_MAX_LIFE_TIME"))
}

func GetPostgresDSN() string {
	return writePostgreDSN(dsnConfig{
		host:     DbHost,
		user:     DbUser,
		password: DbPassword,
		db:       DbName,
		port:     DbPort,
	})
}

func writePostgreDSN(dsn dsnConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dsn.host, dsn.user, dsn.password, dsn.db, dsn.port)
}
