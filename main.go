package main

import (
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
)

func init(){
    logger.EnableLogger()
}

func main() {
    logger.Log.Info("Hello, world!")
}
