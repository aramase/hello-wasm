//go:build tinygo.wasm

package main

import (
	"context"
	"strconv"

	"github.com/aramase/hello-wasm/greeting"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	greeting.RegisterGreeter(&greeter{})
}

type greeter struct{}

func (g greeter) Greet(ctx context.Context, request greeting.GreetRequest) (greeting.GreetReply, error) {
	hostFunctions := greeting.NewHostFunctions()

	shouldGreet, err := hostFunctions.ShouldGreet(ctx, greeting.ShouldGreetRequest{})
	if err != nil {
		return greeting.GreetReply{}, err
	}

	return greeting.GreetReply{
		Message: "Hello, " + request.GetName() + ", Allowed: " + strconv.FormatBool(shouldGreet.GetShouldGreet()),
	}, nil
}
