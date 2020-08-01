package service

import (
	"api/model"
	"errors"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func initDbEngine() *xorm.Engine {
	driverName := "mysql"
	DsName := "root:root@(localhost:3306)/gin?charset=utf8"
	err := errors.New("")
	dbEngine, err := xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	dbEngine.ShowSQL(true)
	dbEngine.SetMaxOpenConns(2)
	dbEngine.SetMapper(core.GonicMapper{})
	err = dbEngine.Sync2(new(model.Task))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("init data base ok")

	return dbEngine
}