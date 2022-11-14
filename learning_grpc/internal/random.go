package internal

import (
	"math/rand"
	"time"
)

func GetRandomTime() (t time.Time) {
	cst := time.FixedZone("CST", 8*60*60)
	year := rand.Intn(200) + 1900
	month := rand.Intn(12)
	day := rand.Intn(28)
	hour := rand.Intn(23)
	min := rand.Intn(59)
	sec := rand.Intn(59)

	return time.Date(year, time.Month(month), day, hour, min, sec, 0, cst)
}
