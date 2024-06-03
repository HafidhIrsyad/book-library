package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func DBConnection() *gorm.DB {
	return makeConnection()
}

func makeConnection() *gorm.DB {
	dsn := GetPostgresDSN()

	fmt.Println(dsn)
	fmt.Println(DriverName)

	db, err := sql.Open(DriverName, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to postgres on makeConnection")
	}

	db.SetMaxOpenConns(DbMaxOpenConnection)
	db.SetMaxIdleConns(DbMaxIdleConnection)
	db.SetConnMaxLifetime(DbConnectionMaxLifeTime)

	config := postgres.Config{
		Conn: db,
	}

	dbGorm, err := gorm.Open(postgres.New(config), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot gorm.Open on makeConnection")
		panic(err)
	}

	return dbGorm
}
