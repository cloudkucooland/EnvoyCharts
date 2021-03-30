package main

import (
	"fmt"
	"github.com/cloudkucooland/EnvoyCharts"
	"time"
)

func main() {
	client, err := envoycharts.New("envoy")
	if err != nil {
		panic(err)
	}

	// catch signals, etc
	ticker := time.Tick(60 * time.Second)
	for range ticker {
		err := client.Sample()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	client.Close()
	fmt.Println("done")
}
