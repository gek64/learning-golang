package dbBase

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var (
	// 用户模拟数据
	users = []User{
		{
			Model:    gorm.Model{ID: 1},
			UserId:   "U1001",
			Name:     "alex",
			Birthday: time.Date(1996, 4, 22, 14, 22, 0, 0, time.UTC),
		},
		{
			Model:    gorm.Model{ID: 2},
			UserId:   "U1002",
			Name:     "bob",
			Birthday: time.Date(1969, 4, 24, 14, 22, 0, 0, time.UTC),
		},
		{
			Model:    gorm.Model{ID: 3},
			UserId:   "U1003",
			Name:     "alice",
			Birthday: time.Date(1696, 4, 26, 14, 22, 0, 0, time.UTC),
		},
	}
	// 产品模拟数据
	products = []Product{
		{
			Model: gorm.Model{ID: 1},
			SKU:   "P1001",
			Name:  "香肠",
			Price: 20,
		},
		{
			Model: gorm.Model{ID: 2},
			SKU:   "P1002",
			Name:  "面包",
			Price: 8,
		},
		{
			Model: gorm.Model{ID: 3},
			SKU:   "P1003",
			Name:  "咖啡",
			Price: 5,
		},
	}
	// 购买记录模拟数据
	purchaseHistories = []PurchaseHistory{
		{
			Model:        gorm.Model{ID: 1},
			PurchaseId:   "H1001",
			SKU:          "P1001",
			UserId:       "U1001",
			PurchaseTime: time.Now().Add(time.Second * 10),
		},
		{
			Model:        gorm.Model{ID: 2},
			PurchaseId:   "H1002",
			SKU:          "P1002",
			UserId:       "U1001",
			PurchaseTime: time.Now().Add(time.Second * 15),
		},
		{
			Model:        gorm.Model{ID: 3},
			PurchaseId:   "H1003",
			SKU:          "P1001",
			UserId:       "U1002",
			PurchaseTime: time.Now().Add(time.Second * 20),
		},
		{
			Model:        gorm.Model{ID: 4},
			PurchaseId:   "H1004",
			SKU:          "P1003",
			UserId:       "U1002",
			PurchaseTime: time.Now().Add(time.Second * 25),
		},
		{
			Model:        gorm.Model{ID: 5},
			PurchaseId:   "H1005",
			SKU:          "P1001",
			UserId:       "U1003",
			PurchaseTime: time.Now().Add(time.Second * 12),
		},
		{
			Model:        gorm.Model{ID: 6},
			PurchaseId:   "H1006",
			SKU:          "P1002",
			UserId:       "U1003",
			PurchaseTime: time.Now().Add(time.Second * 8),
		},
		{
			Model:        gorm.Model{ID: 7},
			PurchaseId:   "H1007",
			SKU:          "P1003",
			UserId:       "U1003",
			PurchaseTime: time.Now().Add(time.Second * 29),
		},
	}
)

// AddSimulatedData 向数据库中插入模拟的数据
func AddSimulatedData(db *gorm.DB) (err error) {
	// 同步表单
	err = db.AutoMigrate(&User{}, &Product{}, &PurchaseHistory{})
	if err != nil {
		return err
	}

	// 配置db数据库源遇到冲突的时候什么都不做,插入模拟数据
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&products)
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&purchaseHistories)

	return nil
}
