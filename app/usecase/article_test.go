package usecase

import (
	"errors"
	"testing"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	mockRepo "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/helper/mock/repository"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestAddAtricles(t *testing.T) {
	mockArticleRepo := &mockRepo.MockArticleRepo{}
	req := &entity.ArticlesListRequest{
		Author: "Test",
		Title:  "Test",
		Body:   "Test",
	}

	res := &entity.Articles{
		Id:        "Suksess",
		Author:    "Suksess",
		Title:     "Suksess",
		Body:      "Suksess",
		CreatedAt: "Suksess",
	}
	Convey("Test Usecase Add Atricles", t, func() {
		Convey("negative scenarios", func() {
			Convey("Request Data nil", func() {
				uc := NewArticleUsecase(mockArticleRepo)
				resp, err := uc.AddAtricles(nil)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("respone err AddAtriclesDB ", func() {
				mockArticleRepo.On("AddAtriclesDB", mock.Anything).Return(nil, errors.New("Some Error")).Once()
				uc := NewArticleUsecase(mockArticleRepo)
				resp, err := uc.AddAtricles(req)
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
		Convey("positive scenarios", func() {
			mockArticleRepo.On("AddAtriclesDB", mock.Anything).Return(res, nil).Once()
			uc := NewArticleUsecase(mockArticleRepo)
			resp, err := uc.AddAtricles(req)
			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
	})
}
