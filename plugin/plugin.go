//go:build tinygo.wasm

package main

import (
	"context"

	"github.com/aramase/hello-wasm/greeting"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	greeting.RegisterGreeter(&greeter{})
}

type greeter struct{}

func (g greeter) Greet(ctx context.Context, request greeting.GreetRequest) (greeting.GreetReply, error) {
	return greeting.GreetReply{
		Message: "Hello, " + request.GetName(),
	}, nil
}
