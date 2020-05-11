
  
package db

import (
	"os"
	// "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm

	"go_api/models"
	"go_api/seed"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	DBPG := "postgres"
	CONNECT := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=disable"
	db, err = gorm.Open(DBPG, CONNECT)
	if err != nil {
		panic(err)
	}

	seed.Load(db)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}