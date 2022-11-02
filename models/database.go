package models

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	var path = fmt.Sprintf("%s: %s@tcp(db:3306)/%s?charest=utf8&parseTime=ture", user, pass, dbName)
	dialector := mysql.Open(path)
	var err error
	if DB, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")
}

func connect(di gorm.Dialector, count uint) {
	var err error
	if DB, err = gorm.Open(di); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(di, count)
			return
		}
		panic(err.Error())
	}
}
