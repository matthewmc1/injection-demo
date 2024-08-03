package main

import (
	"context"
	"fmt"
	"validator/circuit/circuit"
)

var fnCount int

func main() {
	c := context.Background()

	//breakerconfig
	cb := &circuit.BreakerConfig{
		MaxRetries:   3,
		Backoff:      10,
		MaxThreshold: 5,
		URL:          "http://localhost:9000",
	}

	req := cb.Breaker(circuit.Request)
	code, err := req(c, cb.URL)
	fmt.Printf("error code %v and error %v\n", code, err)
	code, err = req(c, "https://google.com")
	fmt.Printf("error code %v and error %v\n", code, err)
}
