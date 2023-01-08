package db

import (
	"errors"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/app/entity"
	"gorm.io/gorm"
)

type ArticlesRepo interface {
	AddAtriclesDB(input *entity.Articles) (interface{}, error)
	GetArticleByParam(req *entity.SearchArticlesRequest) (interface{}, error)
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
	var resRecode []*entity.Articles
	resDb := m.db.Debug().Where("author = ? and title = ?", req.Author, req.Title).First(&resRecode)
	if resDb.Error != nil {
		return nil, resDb.Error
	}
	if resDb.RowsAffected > 1 {
		return nil, errors.New("Data atricle alredy exist")
	}

	var res []*entity.Articles
	err := m.db.Debug().Model(&res).Create(&req).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *articlesRepo) GetArticleByParam(req *entity.SearchArticlesRequest) (interface{}, error) {
	var res []*entity.Articles
	trx := m.db
	if req.Search != "" {
		trx = trx.Where("title like ? or body like ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if req.Author != "" {
		trx = trx.Where("author = ?", req.Author)
	}

	resDb := trx.Debug().Order("author ASC").Find(&res)

	if resDb.Error != nil {
		return nil, resDb.Error
	}
	if len(res) < 1 {
		return nil, errors.New("No recode found")
	}

	return res, nil
}
