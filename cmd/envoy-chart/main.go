package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudkucooland/EnvoyCharts"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var client *envoycharts.Client

func main() {
	var err error
	client, err = envoycharts.New("envoy")
	if err != nil {
		panic(err)
	}

	// start the poller to query the envoy
	go poller()

	// start the http service
	go webservice()

	// listen for a shutdown signal
	sigch := make(chan os.Signal, 3)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)

	sig := <-sigch
	fmt.Println("shutting down", sig)
	client.Close()
}

func webservice() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	barchart(w)
}

func barchart(w io.Writer) {
	e, err := client.GetAll()
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

	bar.SetXAxis(dates).
		AddSeries("Prod", barProd).
		AddSeries("Cons", barCon).
		AddSeries("Net", barNet)
	bar.Render(w)
}

func poller() {
	// catch signals, etc
	ticker := time.Tick(60 * time.Second)
	for range ticker {
		err := client.Sample()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	fmt.Println("shutting down poller")
}
