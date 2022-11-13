package business

import (
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"time"
)

// FindPurchasedRecordBase 查询购买记录会话
// https://gorm.io/zh_CN/docs/method_chaining.html
// 使用新建会话方式创建一个共享的预查询会话(公共的预查询函数都应该返回新建会话方法Session、WithContext、Debug,不然会产生查询之间的互相污染)
func FindPurchasedRecordBase(db *gorm.DB) (tx *gorm.DB) {
	return db.Table("purchase_histories ph").
		Select("ph.user_id, ph.sku, ph.purchase_time, u.name as 'user_name', p.name as 'product_name', p.price").
		Joins("inner join users u on ph.user_id = u.user_id").
		Joins("inner join products p on ph.sku = p.sku").Session(&gorm.Session{})

	//SELECT u.user_id   as 'user_id',
	//	u.name           as 'user',
	//	p.sku            as 'sku',
	//	p.name           as 'product',
	//	ph.purchase_time as 'purchase_time'
	//FROM purchase_histories ph
	//inner join users u on ph.user_id = u.user_id
	//inner join products p on ph.sku = p.sku;
}

// FindPurchasedRecord 查询购买记录
func FindPurchasedRecord(db *gorm.DB) (purchasedResults []PurchasedRecord, err error) {
	// 也可以使用[]map来接收
	//var rawData []map[string]interface{}

	// 查询
	err = FindPurchasedRecordBase(db).Scan(&purchasedResults).Error

	return purchasedResults, err
}

// FindPurchasedRecordByTime 根据时间查询购买记录
func FindPurchasedRecordByTime(db *gorm.DB, startTime time.Time, endTime time.Time) (purchasedResults []PurchasedRecord, err error) {
	// 查询
	err = FindPurchasedRecordBase(db).
		Where("ph.purchase_time >= ? and ph.purchase_time <= ?", startTime, endTime).
		Scan(&purchasedResults).Error

	return purchasedResults, err
}

// FindConsumption 获取消费清单
func FindConsumption(db *gorm.DB) (consumptions []Consumption, err error) {
	records, err := FindPurchasedRecord(db)
	if err != nil {
		return nil, err
	}

	// 筛选符合条件的所有切片条目
	//r := funk.Filter(records,
	//	func(pr PurchasedRecord) bool {
	//		if pr.UserId == userId {
	//			return true
	//		}
	//		return false
	//	})

	// 抽取用户ID列
	r := funk.Get(records, "UserId").([]string)
	// 用户ID去重
	userIds := funk.UniqString(r)

	for _, userId := range userIds {
		var consumption Consumption

		// 按用户ID抽取购买记录的价格到切片
		priceList := funk.FlatMap(records,
			func(pr PurchasedRecord) []float64 {
				if pr.UserId == userId {
					return append([]float64{}, pr.Price)
				}
				return []float64{}
			}).([]float64)
		// 按用户ID抽取购买记录的商品名称到切片
		productList := funk.FlatMap(records,
			func(pr PurchasedRecord) []string {
				if pr.UserId == userId {
					return append([]string{}, pr.ProductName)
				}
				return []string{}
			}).([]string)

		// 返回结构体赋值
		consumption.UserId = userId
		// 循环查询切片记录获取用户ID对应的用户名
		for _, record := range records {
			if record.UserId == userId {
				consumption.Name = record.UserName
				break
			}
		}
		consumption.ProductList = productList
		consumption.Total = funk.SumFloat64(priceList)
		consumptions = append(consumptions, consumption)
	}

	return consumptions, err
}
