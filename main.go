package main

import (
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/db"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/dotenv"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/routes"
)

func init(){
    dotenv.Enable()
    logger.EnableLogger()
    dbConn, err := db.GetDB()
    if err != nil {
      // Handle database connection error (e.g., log and exit)
      panic(err)
    }
    err = routes.Setup(dbConn)
    if err != nil {
        // Handle error setting up routes
        panic(err)
    }
}

func main() {
    // ctx := context.Background()
    logger.Log.Info("Hello, world!")
}
