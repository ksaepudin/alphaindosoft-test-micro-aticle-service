package article

import (
	repo "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/repository/db"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/service"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/usecase"
	"github.com/labstack/echo/v4"
)

func Articles(e *echo.Echo) {

	articlesRepo := repo.NewArticleRepo(repo.ConnArticleDb)
	articlesUsecase := usecase.NewArticleUsecase(articlesRepo)
	articleService := service.NewArticle(articlesUsecase)
	grpArticle := e.Group("/article")
	grpArticle.POST("", articleService.AddAtricles)
	grpArticle.GET("", articleService.GetArticleByParam)
	// grpArticle.GET(":article", articleService.GetArticleByParam)
}
