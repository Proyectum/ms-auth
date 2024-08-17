package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"log"
)

func executeMigrations() {
	dbUser := viper.GetString("data.datasource.postgres.user")
	host := viper.GetString("data.datasource.postgres.host")
	dbPass := viper.GetString("data.datasource.postgres.password")
	db := viper.GetString("data.datasource.postgres.database")
	port := viper.GetInt("data.datasource.postgres.port")

	urlCon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, host, port, db)
	m, err := migrate.New(
		"file://resources/migrations",
		urlCon)

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
		return
	}

	log.Println("Migrations executed successfully")
}
