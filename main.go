package main

import (
	"log"
	"my-gram/config"
	"my-gram/di"
	"my-gram/infrastructure/http/router"
)

func main() {
	config.ReadConfig()

	appController, err := di.NewAppControllerInstance()
	if err != nil {
		log.Fatalln("An error occurred while initializing the app: ", err)
		return
	}

	ServerPort := ":" + config.C.Server.Port
	log.Println("server will run on port: ", ServerPort)

	err = router.NewRouter(appController).Run(ServerPort)
	if err != nil {
		log.Fatalln("An error occurred while running the app: ", err)
		return
	}
}
