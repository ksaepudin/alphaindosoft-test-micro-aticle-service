package db

import (
	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	"gorm.io/gorm"
)

type ArticlesRepo interface {
	AddAtriclesDB(input *entity.Articles) (interface{}, error)
}
type articlesRepo struct {
	db *gorm.DB
}

func NewArticleRepo(conn DbDriver) ArticlesRepo {
	return &articlesRepo{
		db: conn.Db().(*gorm.DB),
	}
}

func (m *articlesRepo) AddAtriclesDB(req *entity.Articles) (interface{}, error) {
	var res []*entity.Articles
	err := m.db.Debug().Model(&res).Create(&req).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
