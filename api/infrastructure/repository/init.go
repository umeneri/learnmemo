package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var dbEngine *xorm.Engine

func init() {
	dbEngine = InitDBEngine()
}

func InitDBEngine() *xorm.Engine {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "root"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "gin"
	}
	dbURL := fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbName)

	var err error
	engine, err := xorm.NewEngine("mysql", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(2)
	engine.SetMapper(core.GonicMapper{})
	_, err = engine.Query("select 1")
	if err != nil {
		log.Fatal("error in db init: ", err)
	}
	log.Println("init data base ok")

	return engine
}
