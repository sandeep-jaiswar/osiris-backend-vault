package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
	"go.uber.org/zap"
)

var (
    dbConn *sql.DB
    once   sync.Once
)

func GetDB() (*sql.DB, error) {
    var err error
	username := os.Getenv("GO_DATABASE_USERNAME")
	password := os.Getenv("GO_DATABASE_PASSWORD")
	hostname := os.Getenv("GO_DATABASE_HOSTNAME")
	port :=os.Getenv("GO_DATABASE_PORT")
	dbname :=os.Getenv("GO_DATABASE_DBNAME")
    once.Do(func() {
        // Initialize the database connection
        dbConn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
        if err != nil {
            logger.Log.Fatal("Error connecting to the database:", zap.Error(err))
            return
        }
        logger.Log.Info("database connected successfully")
    })
    return dbConn, err
}
