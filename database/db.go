package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// for database connected
	_ "github.com/go-sql-driver/mysql"
)

//MysqlDB -
var MysqlDB *gorm.DB

func init() {
	var err error
	MysqlDB, err = gorm.Open("mysql", "k8s_test:k8s_test1234@tcp(10.55.8.211:3306)/k8s_test?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		log.Panic(err)
	}
	MysqlDB.DB().SetMaxIdleConns(3)
	MysqlDB.DB().SetMaxOpenConns(10)
	MysqlDB.DB().SetConnMaxLifetime(time.Minute * 5)

}
