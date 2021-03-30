package main

import (
	"github.com/cloudkucooland/EnvoyCharts"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
	"time"
)

func main() {
	c, err := envoycharts.New()
	if err != nil {
		panic(err)
	}

	e, err := c.GetAll()
	if err != nil {
		panic(err)
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Solar Production"}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
	)

	barProd := make([]opts.BarData, 0)
	barCon := make([]opts.BarData, 0)
	barNet := make([]opts.BarData, 0)
	dates := make([]string, 0)

	for _, v := range e {
		fu := time.Unix(v.Date, 0)
		t := fu.Format("2006/01/02 15:04:05")

		dates = append(dates, t)
		barProd = append(barProd, opts.BarData{Value: v.ProductionW})
		barCon = append(barCon, opts.BarData{Value: 0 - v.ConsumptionW})
		barNet = append(barNet, opts.BarData{Value: 0 - v.NetW})
	}

	f, err := os.Create("bar.html")
	if err != nil {
		panic(err)
	}
	bar.SetXAxis(dates).
		AddSeries("Prod", barProd).
		AddSeries("Cons", barCon).
		AddSeries("Net", barNet)
	bar.Render(f)
}
