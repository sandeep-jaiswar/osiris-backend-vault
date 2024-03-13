package main

import (
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/db"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/dotenv"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
)

func init(){
    dotenv.Enable()
    logger.EnableLogger()
    db.GetDB()
}

func main() {
    logger.Log.Info("Hello, world!")
}
