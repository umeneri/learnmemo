package service

import (
	"errors"
	"fmt"
	"log"
	"api/model"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:root@(localhost:3306)/gin?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	err = DbEngine.Sync2(new(model.Book))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("init data base ok")
}
