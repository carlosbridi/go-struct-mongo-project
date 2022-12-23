package domains

import (
	"time"
)

type Invoice struct {
	Id         string
	Customer   Customer
	Value      float64
	OrderDate  time.Time
	CreateDate time.Time
}
