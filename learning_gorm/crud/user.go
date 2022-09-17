package crud

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func NewUserDB() (db *gorm.DB, err error) {
	// 打开数据库
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}

	// 同步表单
	err = db.AutoMigrate(&User{})
	if err != nil {
		return db, err
	}

	return db, nil
}
