package repository

import (
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepo struct {
	mock.Mock
}

func (m *MockArticleRepo) AddAtriclesDB(req *entity.Articles) (interface{}, error) {
	call := m.Called(req)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}

func (m *MockArticleRepo) GetArticleByParam(req *entity.SearchArticlesRequest) (interface{}, error) {
	call := m.Called(req)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return &res, call.Error(1)
}
