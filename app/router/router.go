package router

import (
	"fmt"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/router/article"
	cfg "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/config"
	"github.com/labstack/echo/v4"
)

func Route() {
	e := echo.New()
	article.Articles(e)
	port := fmt.Sprintf(":%s", cfg.GetConfig().Service.Articles.Port)
	e.Logger.Fatal(e.Start(port))
}
