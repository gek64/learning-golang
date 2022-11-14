package internal

import (
	"time"
)

func GetRandomTime() (t time.Time) {
	cst := time.FixedZone("CST", 8*60*60)
	newTime := time.Date(1996, 4, 22, 16, 22, 00, 00, cst)
	return newTime
}
