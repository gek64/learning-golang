package dbBase

import (
	"gorm.io/gorm"
	"time"
)

// User 用户表
type User struct {
	gorm.Model
	UserId   string `gorm:"uniqueIndex;not null"`
	Name     string
	Birthday time.Time
}

// Product 产品表
type Product struct {
	gorm.Model
	SKU   string `gorm:"uniqueIndex;not null"`
	Name  string
	Price int
}

// PurchaseHistory 购买记录表
type PurchaseHistory struct {
	gorm.Model
	PurchaseId   string `gorm:"uniqueIndex;not null"`
	SKU          string `gorm:"index;not null"`
	UserId       string `gorm:"index;not null"`
	PurchaseTime time.Time
}
