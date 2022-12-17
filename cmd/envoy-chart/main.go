package main

import (
	"context"
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
var pollrate time.Duration

func main() {
	var err error
    pollrate = 5

	ctx, shutdownpoller := context.WithCancel(context.Background())

	// "" triggers dns-sd discovery
	client, err = envoycharts.New("")
	if err != nil {
		panic(err)
	}

	// start the poller to query the envoy
	go poller(ctx)

	// start the http service
	srv := &http.Server{Addr: ":8081"}
	go webservice(srv)

	// listen for a shutdown signal
	sigch := make(chan os.Signal, 3)
	signal.Notify(sigch, syscall.SIGQUIT, syscall.SIGTERM)
	<-sigch // wait here until an OS signal is sent

	fmt.Println("shutting down")

	// shutdown the webservice
	webshutdown, cancelwebshutdown := context.WithTimeout(ctx, 5*time.Second)
	if err := srv.Shutdown(webshutdown); err != nil {
		fmt.Println(err.Error())
		cancelwebshutdown()
	} else {
		<-webshutdown.Done()
	}

	// now shutdown the poller
	shutdownpoller()
	// final cleanup
	client.Close()
}

// the IP address of the envoy can change over time,
// this tries to re-discover and keep polling if that happens
func poller(ctx context.Context) {
	defer func() {
		fmt.Println("shutting down poller")
	}()

	ticker := time.NewTicker(pollrate * time.Minute)
	defer ticker.Stop()

	err := client.Sample()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		// fmt.Println("loop tick")
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
            if !client.Configured() {
				fmt.Println("not configuring, trying again")
				client.Reset()
                continue
            }

			// fmt.Println("starting sample")
			if err = client.Sample(); err != nil {
				fmt.Println(err.Error())
				fmt.Println("reseting client")
                client.Reset()
                continue
			}
			// fmt.Println("sample complete")
		}
	}
}

func webservice(srv *http.Server) {
	sm := http.NewServeMux()
	sm.HandleFunc("/", handler)

	srv.Handler = sm
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	chart, ok := r.Form["c"]
	if ok && chart[0] == "daily" {
		dailyChart(w)
		return
	}

	date, ok := r.Form["d"]
	if !ok {
		// start day not set, just show the past 24 hours
		pastDay(w)
		return
	}
	start, err := time.ParseInLocation("2006/01/02", date[0], client.TZ)
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
	end, err := time.ParseInLocation("2006/01/02", enddate[0], client.TZ)
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
		fmt.Println(err.Error())
		return
	}
	title := fmt.Sprintf("Solar Production for %s", t.Format("2006/01/02"))
	envoycharts.Linechart(w, samples, title, client.TZ)
}

func dayRange(w io.Writer, start, end time.Time) {
	samples, err := client.GetDayRange(start, end)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	title := fmt.Sprintf("Solar Production %s - %s", start.Format("2006/01/02"), end.Format("2006/01/02"))
	envoycharts.Linechart(w, samples, title, client.TZ)
}

func pastDay(w io.Writer) {
	samples, err := client.GetPastDay()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	envoycharts.Linechart(w, samples, "Solar Production for Past 24 hours", client.TZ)
}

func dailyChart(w io.Writer) {
	ds, err := client.GetAllDaily()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	envoycharts.LinechartDaily(w, ds, "Daily Totals", client.TZ)
}
