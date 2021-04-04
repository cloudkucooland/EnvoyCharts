package model

import (
	"time"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

// Entry is an individual system sample, the basic unit stored in the database

// next schema change, convert Date to time.Time `objectbox:"date"`
type Entry struct {
	Id           int64 `objectbox:"id"`
	Date         int64
	ProductionW  float64
	ConsumptionW float64
	NetW         float64
}

type Daily struct {
	DID            int64     `objectbox:"id(assignable)","unique"`
	Date           time.Time `objectbox:"date"`
	ProductionkWh  float64
	ConsumptionkWh float64
}
