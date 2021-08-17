package main

import (
	"github.com/edvaldo-domingos/go_banking/app"
	"github.com/edvaldo-domingos/go_banking/logger"
)


func main(){
	// log.Println("stating application")
	logger.Info("stating application")
	app.Start()
}

