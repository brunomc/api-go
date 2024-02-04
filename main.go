package main

import (
	"github.com/brunomc/api-go/config"
	"github.com/brunomc/api-go/router"
)

var (
	logger *config.Logger
)

func main() {
	err := config.Init()
	logger = config.GetLogger("main")
	if err != nil {
		//panic(err)
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
