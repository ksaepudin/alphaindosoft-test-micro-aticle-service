package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	repo "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/repository/db"
)

type ArticlesUsecase interface {
	AddAtricles(data *entity.ArticlesListRequest) (interface{}, error)
}

type articles struct {
	repoArticle repo.ArticlesRepo
}

func NewArticleUsecase(repoArticle repo.ArticlesRepo) ArticlesUsecase {
	return &articles{repoArticle: repoArticle}
}

func (m *articles) AddAtricles(data *entity.ArticlesListRequest) (interface{}, error) {
	uuid := uuid.New().String()
	req := &entity.Articles{
		Id:        uuid,
		Author:    data.Author,
		Title:     data.Title,
		Body:      data.Body,
		CreatedAt: time.Now().String(),
	}
	return m.repoArticle.AddAtriclesDB(req)
	//  data, nil
}
