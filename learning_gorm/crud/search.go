package crud

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// SearchPrimaryKey 查询单个数据,以主键/第一个字段作为排序和条件的依据
func SearchPrimaryKey() (err error) {

	db, err := NewUserDB()
	if err != nil {
		return err
	}

	// 创建一个记录
	user := User{}

	// 获取第一条记录（主键升序）
	result := db.First(&user, "name = ?", "bob")
	// 返回找到的记录数
	//result.RowsAffected
	// returns error or nil
	//result.Error

	// 获取一条记录，没有指定排序字段
	db.Take(&user, "name = ?", "bob")

	// 获取最后一条记录（主键降序）
	db.Last(&user, "name = ?", "bob")

	// 检查 ErrRecordNotFound 错误
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
	fmt.Println(user, result.RowsAffected, result.Error)

	return nil
}

// SearchAll 查询多个数据
func SearchAll() (err error) {

	db, err := NewUserDB()
	if err != nil {
		return err
	}

	// 创建一个记录
	var users []User

	result := db.Find(&users, "name = ?", "bob")

	// 检查 ErrRecordNotFound 错误
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
	fmt.Println(users, result.RowsAffected, result.Error)

	return nil
}
