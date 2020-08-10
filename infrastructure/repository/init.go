package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func initDbEngine(dbName string) *xorm.Engine {
	dbURL := os.Getenv("DATABASE_URL")
	driverName := "mysql"
	if dbURL == "" {
		dbURL = fmt.Sprintf("root:root@(localhost:3306)/%s?charset=utf8", dbName)
	}
	err := errors.New("")
	dbEngine, err := xorm.NewEngine(driverName, dbURL)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	dbEngine.ShowSQL(true)
	dbEngine.SetMaxOpenConns(2)
	dbEngine.SetMapper(core.GonicMapper{})
	fmt.Println("init data base ok")

	return dbEngine
}
