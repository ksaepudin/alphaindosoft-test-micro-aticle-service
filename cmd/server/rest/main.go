package main

import (
	"fmt"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/router"
	cfg "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/config"
)

func main() {
	config := cfg.GetConfig()
	fmt.Println(config)
	router.Route()
}
