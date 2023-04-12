package main

import (
	"log"
	"my-gram/config"
	"my-gram/di"
	"my-gram/infrastructure/http/router"
)

// main Function to run the app
// @Title My Gram API Documentation
// @Description My Gram API is a final project of Hacktiv8 Golang Bootcamp. This API is used to manage social media accounts and photos.
// @Version 1.0.0
// @contact.name Fikri Yusrihan
// @contact.url https://fikriyusrihan.github.io
// @contact.email fikriyusrihan@gmail.com
// @license MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8000
// @BasePath /
// @securityDefinitions.jwt Bearer
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
