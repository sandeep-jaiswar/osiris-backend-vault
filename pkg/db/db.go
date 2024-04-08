package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func GetDB() (*gorm.DB, error) {
	var err error

	username := os.Getenv("GO_DATABASE_USERNAME")
	password := os.Getenv("GO_DATABASE_PASSWORD")
	hostname := os.Getenv("GO_DATABASE_HOSTNAME")
	port := os.Getenv("GO_DATABASE_PORT")
	dbname := os.Getenv("GO_DATABASE_DBNAME")

	once.Do(func() {
		// Initialize the GORM database connection
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Log.Fatal("Error connecting to the database:", zap.Error(err))
			return
		}
		logger.Log.Info("Database connected successfully")
	})

	return dbConn, err
}
