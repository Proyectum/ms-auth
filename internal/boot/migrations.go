package boot

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func ExecuteMigrations() {
	database := CONFIG.Data.Datasource.Postgres

	urlCon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", database.User, database.Password,
		database.Host, database.Port, database.Database)

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
