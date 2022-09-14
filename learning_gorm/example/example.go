package example

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	// 继承 gorm.Model 结构体
	gorm.Model
	Code  string
	Price uint
}

func RunExample() {
	// 打开数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("failed to connect database")
	}

	// 迁移 schema
	// 同步表结构到数据库
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Panicln(err)
	}

	// 增,创建一个实例
	db.Create(&Product{
		Code:  "1C6500",
		Price: 43,
	})

	// 查,查询实例到变量
	var p1 Product
	db.Find(&p1, "code = ?", "1C6500")

	// 改,更新实例
	db.Model(&p1).Update("price", 200)
	db.Model(&p1).Updates(Product{
		Code:  "1C6501",
		Price: 100,
	})
	db.Model(&p1).Updates(map[string]interface{}{
		"Code":  "1C6501",
		"Price": 100,
	})

	// 删除
	db.Delete(&p1, 1)
}
