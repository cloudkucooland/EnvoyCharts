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
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	date, ok := r.Form["d"]
	if !ok {
		// start day not set, just show the past 24 hours
		pastDay(w)
		return
	}
	start, err := time.Parse("2006/01/02", date[0])
	if err != nil {
		fmt.Println(err)
		// start day not valid, show past 24h
		pastDay(w)
		return
	}

	enddate, ok := r.Form["end"]
	if !ok {
		// start valid, but end not set
		specificDay(w, start)
		return
	}
	end, err := time.Parse("2006/01/02", enddate[0])
	if err != nil {
		fmt.Println(err)
		// start valid, but end invalid
		specificDay(w, start)
		return
	}

	// both start and end set and valid, show the range
	dayRange(w, start, end)
}

func specificDay(w io.Writer, t time.Time) {
	samples, err := client.GetDay(t)
	if err != nil {
		panic(err)
	}
	title := fmt.Sprintf("Solar Production for %s", t.Format("2006/01/02"))
	envoycharts.Linechart(w, samples, title)
}

func dayRange(w io.Writer, start, end time.Time) {
	samples, err := client.GetDayRange(start, end)
	if err != nil {
		panic(err)
	}
	title := fmt.Sprintf("Solar Production %s - %s", start.Format("2006/01/02"), end.Format("2006/01/02"))
	envoycharts.Linechart(w, samples, title)
}

func pastDay(w io.Writer) {
	samples, err := client.GetPastDay()
	if err != nil {
		panic(err)
	}
	envoycharts.Linechart(w, samples, "Solar Production for Past 24 hours")
}

func poller() {
	ticker := time.Tick(600 * time.Second)

	for range ticker {
		err := client.Sample()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	fmt.Println("shutting down poller")
}
