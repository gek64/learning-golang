package main

import (
	"fmt"
	"learning_gorm/business"
	"learning_gorm/dbBase"
	"learning_gorm/dbOperater"
	"log"
	"time"
)

func main() {
	// 连接到数据库
	db, err := dbOperater.NewUserMysqlDB()
	if err != nil {
		log.Panicln(err)
	}
	// 数据库建立表单和插入模拟数据
	err = dbBase.AddSimulatedData(db)
	if err != nil {
		log.Panicln(err)
	}

	startTime := time.Date(2020, 11, 07, 9, 28, 00, 00, time.FixedZone("CST", 8*60*60))
	endTime := time.Date(2030, 11, 07, 9, 29, 00, 00, time.FixedZone("CST", 8*60*60))
	// 查询时间段内的购买记录
	records, err := business.FindPurchasedRecordByTime(db, startTime, endTime)
	if err != nil {
		log.Panicln(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}
	// 查询每一个用户的购物清单
	consumptions, err := business.FindConsumption(db)
	if err != nil {
		log.Panicln(err)
	}
	for _, consumption := range consumptions {
		fmt.Println(consumption)
	}
}
