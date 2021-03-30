package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

// Entry is an individual system sample, the basic unit stored in the database

// next schema change, convert Date to time.Time `objectbox:"date"`
type Entry struct {
	Id           int64
	Date         int64
	ProductionW  float64
	ConsumptionW float64
	NetW         float64
}
