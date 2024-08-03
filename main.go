package main

// A circuit breaker is a system that monitors for failures in a function. If the function
// has failed too many times in a row in the past, the circuit breaker goes in fail-fast
// mode which stops any invocations in the future, instead it quickly returns an error.
// This is often used to prevent cascading failures to external services.

// Create a circuit breaker that accepts a function and a consecutive failure threshold (int).

// This exercise can be completed in any language, below we provide some example signatures
// in popular languages. Feel free to work in whichever editor/IDE you prefer.

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"validator/circuit/circuit"
)

var fnCount int

func makeRequest(context.Context) (string, error) {
	fmt.Println("hello")
	return "request-service", errors.New("Request")
}

func Breaker(circuit func(context.Context) (string, error), failureThreshold uint) func(context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		if failureThreshold == 0 {
			fmt.Println("invalid threshold")
			return "Error with threshold value", errors.New("invalid threshold value")
		}

		fnCount++
		slog.Info("function invocation", "count", fnCount)
		if fnCount < int(failureThreshold) {
			fmt.Println("under request threshold, request is allowed")
			return "request ok", nil
		} else {
			fmt.Println("equal to or over request threshold")
			return "request blocked", errors.New("threshold exceeded")
		}
	}
}

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
