package service

import (
	"errors"
	"net/http"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/helper"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/usecase"
	"github.com/labstack/echo/v4"
)

type ArticlesService interface {
	AddAtricles(c echo.Context) error
}

type articles struct {
	uc usecase.ArticlesUsecase
}

func NewArticle(uc usecase.ArticlesUsecase) ArticlesService {
	return &articles{uc: uc}
}

func (m *articles) AddAtricles(c echo.Context) error {
	req := new(entity.ArticlesListRequest)
	if err := c.Bind(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, errors.New("binding element must be a struct").Error())
	}

	res, err := m.uc.AddAtricles(req)
	if err != nil {
		return helper.Response(c, http.StatusFailedDependency, "Usecase Error", err)
	}
	return helper.Response(c, http.StatusOK, "Yey Berhasil", res)
}
