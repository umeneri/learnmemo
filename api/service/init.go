package service

import (
	"api/model"
	"errors"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
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
	DbEngine.SetMapper(core.GonicMapper{})
	err = DbEngine.Sync2(new(model.Task))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("init data base ok")
}
