package envoycharts

import (
	"time"
)

type Sample struct {
	Id           int64
	Date         time.Time
	ProductionW  float64
	ConsumptionW float64
	NetW         float64
}

type Daily struct {
	Id             int64
	Date           time.Time
	ProductionkWh  float64
	ConsumptionkWh float64
}
