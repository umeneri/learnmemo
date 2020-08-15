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
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "gin"
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = fmt.Sprintf("root:root@(localhost:3306)/%s?charset=utf8", dbName)
	}
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
