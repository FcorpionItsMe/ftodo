package main

import (
	"github.com/FcorpionItsMe/ftodo/internal/app"
	"github.com/FcorpionItsMe/ftodo/internal/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//For develop
	setEnvErr := os.Setenv("FTODO_CONFIG_PATH", "configs/local.yaml")
	if setEnvErr != nil {
		log.Println(setEnvErr)
		return
	}
	godotenv.Load(".env")
	//For develop

	cfg := config.MustLoad()
	app.Run(cfg)
}
