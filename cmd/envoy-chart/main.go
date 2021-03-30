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

func handler(w http.ResponseWriter, _ *http.Request) {
	pastDay(w)
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
