package db

import (
	"log"

	cfg "github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/config"
)

var (
	ConnArticleDb DbDriver
)

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
}

func init() {
	var err error
	ConnArticleDb, err = NewMySQLDriver(cfg.GetConfig().Database.Articles.Mysql)
	if err != nil {
		log.Fatalf("unable to connect to Article DB")
	}
}
