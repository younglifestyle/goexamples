package db

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DBPool struct {
	DbFiles *gorm.DB
}

var (
	// mysql conn
	dbp DBPool
)

func Con() DBPool {
	return dbp
}

func InitMysqlDb() {
	var p *sql.DB

	dsn := viper.GetString("dsn")
	if dsn == "" {
		log.Fatal("please add dsn in the ./cfgfile/cfg.json")
	}

	portal, err := gorm.Open("mysql", dsn)
	portal.Dialect().SetDB(p)
	if err != nil {
		log.Fatal("connect to falcon_portal: %s", err.Error())
	}

	// test db is true connect
	err = portal.DB().Ping()
	if err != nil {
		log.Panic("failed connect database : %v", err)
	}
	portal.DB().SetMaxIdleConns(20)

	portal.SingularTable(true)
	dbp.DbFiles = portal

	return
}
