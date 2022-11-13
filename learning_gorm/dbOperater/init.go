package dbOperater

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewUserMysqlDB() (db *gorm.DB, err error) {
	// 自定义mysql数据源名称
	dsn := mysql.Config{
		// DSN data source name
		DSN: "root:admin@tcp(192.168.137.193:3306)/gorm?charset=utf8&parseTime=True&loc=Local",
		// string 类型字段的默认长度
		DefaultStringSize: 256,
		// 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DisableDatetimePrecision: true,
		// 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameIndex: true,
		// 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		DontSupportRenameColumn: true,
		// 根据当前 MySQL 版本自动配置
		SkipInitializeWithVersion: false,
	}

	// 返回数据库
	return gorm.Open(mysql.New(dsn), &gorm.Config{})
}
