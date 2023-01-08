package service

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	mockRepo "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/helper/mock/repository"
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/usecase"
	"github.com/labstack/echo/v4"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/stretchr/testify/mock"
)

func TestAddAtricles(t *testing.T) {
	var (
		mockDB = []*entity.Articles{
			&entity.Articles{Id: "test", Author: "test", Title: "test", Body: "test", CreatedAt: "test"},
		}
		articleJSONSuccsess = `{
			"author": "Test",
			"title": "Test",
			"body": "Test"
		}`
	)
	mockArticleRepo := &mockRepo.MockArticleRepo{}
	e := echo.New()
	Convey("Test Service Add Atricles", t, func() {
		Convey("Positive Scenario", func() {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(articleJSONSuccsess))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockArticleRepo.On("AddAtriclesDB", mock.Anything).Return(mockDB, nil).Once()
			uc := usecase.NewArticleUsecase(mockArticleRepo)
			svc := NewArticle(uc)
			err := svc.AddAtricles(c)
			So(err, ShouldBeNil)
		})
		Convey("Negative Scenario", func() {
			Convey("Reuqest Failed Scenario", func() {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"failed"}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				uc := usecase.NewArticleUsecase(mockArticleRepo)
				svc := NewArticle(uc)
				err := svc.AddAtricles(c)
				So(err, ShouldBeNil)
			})
			Convey("Failed AddArticle Scenario", func() {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(articleJSONSuccsess))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				mockArticleRepo.On("AddAtriclesDB", mock.Anything).Return(nil, errors.New("some error")).Once()
				uc := usecase.NewArticleUsecase(mockArticleRepo)
				svc := NewArticle(uc)
				err := svc.AddAtricles(c)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestGetArticleByParam(t *testing.T) {
	mockArticleRepo := &mockRepo.MockArticleRepo{}
	var (
		mockDB = []*entity.Articles{
			{Id: "test", Author: "test", Title: "test", Body: "test", CreatedAt: "test"},
		}
	)
	e := echo.New()
	Convey("Test Service GetArticleByParam", t, func() {
		Convey("Positive Scenario", func() {
			Convey("Positive By Params Scenario", func() {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/")
				c.SetParamNames("search", "author")
				c.SetParamValues("test", "test")
				mockArticleRepo.On("GetArticleByParam", mock.Anything).Return(mockDB, nil).Once()
				uc := usecase.NewArticleUsecase(mockArticleRepo)
				svc := NewArticle(uc)
				err := svc.GetArticleByParam(c)
				So(err, ShouldBeNil)
			})
			Convey("Positive By GetAll Scenario", func() {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				mockArticleRepo.On("GetArticleByParam", mock.Anything).Return(mockDB, nil).Once()
				uc := usecase.NewArticleUsecase(mockArticleRepo)
				svc := NewArticle(uc)
				err := svc.GetArticleByParam(c)
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("Reuqest Failed Scenario", func() {
				req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"failed"}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/")
				c.SetParamNames("worng")
				c.SetParamValues("type")
				uc := usecase.NewArticleUsecase(mockArticleRepo)
				svc := NewArticle(uc)
				err := svc.GetArticleByParam(c)
				So(err, ShouldBeNil)
			})
			Convey("Failed GetArticleByParam Scenario", func() {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/")
				c.SetParamNames("search", "author")
				c.SetParamValues("test", "test")
				mockArticleRepo.On("GetArticleByParam", mock.Anything).Return(nil, errors.New("Some Error")).Once()
				uc := usecase.NewArticleUsecase(mockArticleRepo)
				svc := NewArticle(uc)
				err := svc.GetArticleByParam(c)
				So(err, ShouldBeNil)
			})
		})
	})
}
