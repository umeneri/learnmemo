package repository

import (
	"api/domain/model"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func initDbEngine(dbName string) *xorm.Engine {
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		mysqlHost = "localhost"
	}

	driverName := "mysql"
	DsName := fmt.Sprintf("root:root@(%s:3306)/%s?charset=utf8", mysqlHost, dbName)
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
