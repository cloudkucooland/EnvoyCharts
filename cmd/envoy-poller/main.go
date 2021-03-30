package main

import (
	"fmt"
	"github.com/cloudkucooland/EnvoyCharts"
	"github.com/cloudkucooland/go-envoy"
	"time"
)

func main() {
	client, err := envoycharts.New()
	if err != nil {
		panic(err)
	}

	envoy, err := envoy.New("envoy")
	if err != nil {
		panic(err)
	}

	// catch signals, etc
	ticker := time.Tick(60 * time.Second)
	for range ticker {
		p, c, n, err := envoy.Now()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		client.Insert(p, c, n)
	}

	fmt.Println("done")
}
