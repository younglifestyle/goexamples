package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	//gorm.Model
	Id        int
	Code      string
	Price     uint
	CreatedAt time.Time  `gorm:"times"`
	DeletedAt *time.Time `sql:"index"`
}
type P1 struct {
	//gorm.Model
	Id   string
	Code string `gorm:"column:code"`
	Test string `gorm:"-"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	//// 自动迁移模式
	if err := db.AutoMigrate(&Product{}).Error; err != nil {
		fmt.Println("error :", err)
		return
	}

	//
	// 创建
	p1 := Product{Code: "L1234242", Price: 1100}
	db.Create(&p1)
	db.Create(&Product{Code: "L1211", Price: 1001})
	e := db.Table("products").Create(&P1{Id: "10", Code: "L1251"}).Error
	fmt.Println("error :", e)

	var product1 Product
	//rows, err := db.DB().Query("SELECT * FROM products WHERE " +
	//	"price=(select max(price) from products)")
	//defer rows.Close()
	//if err != nil {
	//	fmt.Println("Query Data is failed", err.Error())
	//	return
	//}
	//
	//for rows.Next() {
	//	rows.Scan(&product1.Id, &product1.Code, &product1.Price)
	//	fmt.Println(product1)
	//}

	db.Create(&Product{Id: 16, Code: "L1213", Price: 12144})
	db.Create(&Product{Id: 134, Code: "L1213", Price: 12144})

	db.Table("products").Where("id = ?", 16).
		Update(&Product{Id: 16, Code: "L1213333", Price: 12144})

	p3 := P1{Code: "21321241214"}
	db.Table("products").Create(&p3)

	fmt.Println("p3 :", p3)

	db.Table("products").
		//Where("id = ?", 16).
		Find(&product1, "id = ?", 16).
		Where("price=(select max(price) from products)").Update("code=?", "L1265")
	fmt.Println("err1 :", product1)

	//db.Where("price=?").Find(&product1)
	fmt.Println("product1 :", product1)

	// 读取
	var p2 P1
	db.Table("products").Select("id,code").
		Where("id = ?", 14).Find(&p2)
	//db.First(&p2, 1) // 查询id为1的product
	fmt.Println("asd", p2)

	//db.First(&p2, 1) // 查询id为1的product

	// 使用first和find查找单个数据时，没有找到会报错，若接受单位为数组则没有数据将不会报错
	var p123 []Product
	d1 := db.Table("products").
		Find(&p123, "code = ?", "L2137421")
	fmt.Println("aids :", d1.Error)

	var ps []Product
	db.Table("products").Find(&ps)
	fmt.Println("asd1", ps)

	var ps1 []Product
	db.Table("products").Offset(1). /*.Find(&ps1)*/ Limit(3).Find(&ps1)
	fmt.Println("test :", ps1)

	db.Table("products").Where("id = ?", 1).Update(&Product{Code: "L1265", Price: 1})

	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//

	db.Create(&Product{Id: 20, Code: "L1213", Price: 12144})

	var test Product
	// 删除 - 删除product  查找数据后删除
	dt := db.Table("products").Where("id = ?", 20).
		First(&test).Delete(nil)
	fmt.Println("error1 :", dt.Error, test)

	var num int
	var testArr []Product
	db.Table("products").Where("code LIKE ?", "L126%").
		Select("id").Group("id").Scan(&testArr).Count(&num)
	fmt.Println("array :", testArr)
	fmt.Println("num :", num)

	// 分开进行编写
	//ds1 := db.Table("products")
	//ds1 = ds1.Group("id")
	//ds1 = ds1.Limit(3)
	//ds1.Scan(&testArr).Count(&num)
	//fmt.Println("array112 :", testArr)

	var testArr1 []Product
	dts := db.Table("products").
		Limit(5).
		//Where("id = ?", "1").Select("id").
		Find(&testArr1)
	dts.Count(&num)
	fmt.Println("num1 :", num, testArr1)

	query := db.Table("products").
		Select("id").
		Where("code in (?)", []string{"L1213", "L1251"}).SubQuery()

	db.Table("products").Select("*").
		Joins("INNER JOIN ? AS t1 using(id)", query).
		Find(&testArr1)
	fmt.Println("num2222 :", testArr1)

	// unsupported destination, should be slice or struct
	//fmt.Println("err01 ,",
	//	db.Table("products").First(nil, "id = ?", 30).Error)

	var res1 Product
	db.Where("id = ?", 1243).Assign(Product{
		Code:  "L12345",
		Price: 123456,
	}).FirstOrInit(&res1)
	fmt.Println("test :", res1)

	// CreatedAt 若指定时间，将按照指定时间存储
	now := time.Date(2018, time.November, 11, 23, 0, 0, 0, time.UTC)
	fmt.Println("time now :", now)
	p21 := Product{Id: 1, Code: "L1234242", Price: 1100, CreatedAt: now}
	db.Create(&p21)
}
