package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	type User struct {
		Info *StringMap `gorm:"column:info;type:json"`
	}

	u := &User{
		Info: NewEmptyStringMap(), // 这样会初始化map，即使不设置任何值，保存后在db中的值为`{}`
	}
	u.Info.Src["nihao"] = "123"

	db, err := gorm.Open("mysql", "root:root@/test_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("error :", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	err = db.Table("users").Create(u).Error
	if err != nil {
		log.Println("error :", err)
		return
	}
}
