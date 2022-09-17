package crud

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// BeforeCreate 钩子 在表单对记录进行操作(BeforeSave, BeforeCreate, AfterSave, AfterCreate)时调用的函数
func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	return nil
}

func Insert() (err error) {
	db, err := NewUserDB()
	if err != nil {
		return err
	}

	// 创建一个记录
	user := User{
		Name:     "lily",
		Age:      22,
		Birthday: time.Now(),
	}

	// 创建记录并更新全部的字段
	db.Create(&user)
	// 创建记录,并更新指定的字段
	//db.Select("Name", "Age").Create(&user)
	// 创建记录,并忽略指定的字段
	//db.Omit("Name", "Age").Create(&user)

	return nil
}

func InsertAll() (err error) {
	// 打开数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// 同步表单
	err = db.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	// 创建多个记录
	users := []User{{
		Name:     "lily",
		Age:      12,
		Birthday: time.Now(),
	}, {
		Name:     "tom",
		Age:      19,
		Birthday: time.Now(),
	}, {
		Name:     "bob",
		Age:      22,
		Birthday: time.Now(),
	}}

	// 创建记录并更新全部的字段
	db.Create(&users)
	// 创建记录并更新全部的字段,跳过钩子函数
	//db.Session(&gorm.Session{SkipHooks: true}).Create(&users)

	// 创建记录,并更新指定的字段
	//db.Select("Name", "Age").Create(&users)
	// 创建记录,并忽略指定的字段
	//db.Omit("Name", "Age").Create(&users)

	return nil
}
