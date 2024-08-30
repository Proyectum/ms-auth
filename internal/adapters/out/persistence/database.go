package persistence

import (
	"fmt"
	"github.com/proyectum/ms-auth/internal/boot"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func getDatasource() *gorm.DB {

	dbOnce.Do(func() {
		databaseConf := boot.CONFIG.Data.Datasource.Postgres

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			databaseConf.Host, databaseConf.User, databaseConf.Password, databaseConf.Database, databaseConf.Port)
		var err error
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect database %w", err))
		}
	})

	return dbInstance
}
