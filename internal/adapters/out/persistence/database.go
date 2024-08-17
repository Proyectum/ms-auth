package persistence

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func getDatasource() *gorm.DB {

	dbOnce.Do(func() {
		dbUser := viper.GetString("data.datasource.postgres.user")
		host := viper.GetString("data.datasource.postgres.host")
		dbPass := viper.GetString("data.datasource.postgres.password")
		dbName := viper.GetString("data.datasource.postgres.database")
		port := viper.GetInt("data.datasource.postgres.port")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", host, dbUser, dbPass, dbName, port)
		var err error
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database")
		}
	})

	return dbInstance
}
