package db

import (
	"fmt"
	"golang_graphql/config"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDatabase() {
	databseURL := config.GetConfig().Database.URL
	migrateConnection, err := migrate.New("file://db/migrate", databseURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	version, _, _ := migrateConnection.Version()
	fmt.Println(version)
	if version != 1 {
		migrateConnection.Migrate(1)
	}
	migrateConnection.Close()

	dbConnection, err := gorm.Open("postgres", databseURL)

	if err != nil {
		fmt.Println(err)
		return
	}
	db = dbConnection
	db.LogMode(config.GetConfig().Database.Logmode)
}

func GetDB() *gorm.DB {
	return db
}
