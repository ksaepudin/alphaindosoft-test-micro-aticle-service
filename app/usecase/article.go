package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	repo "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/repository/db"
)

type ArticlesUsecase interface {
	AddAtricles(data interface{}) (interface{}, error)
}

type articles struct {
	repoArticle repo.ArticlesRepo
}

func NewArticleUsecase(repoArticle repo.ArticlesRepo) ArticlesUsecase {
	return &articles{repoArticle: repoArticle}
}

func (m *articles) AddAtricles(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("Requset Data Nil")
	}
	reqInput := data.(*entity.ArticlesListRequest)
	uuid := uuid.New().String()
	req := &entity.Articles{
		Id:        uuid,
		Author:    reqInput.Author,
		Title:     reqInput.Title,
		Body:      reqInput.Body,
		CreatedAt: time.Now().String(),
	}

	res, err := m.repoArticle.AddAtriclesDB(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
