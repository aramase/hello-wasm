package main

import (
	"context"
	"fmt"

	"github.com/aramase/hello-wasm/greeting"
)

func main() {
	ctx := context.Background()
	p, err := greeting.NewGreeterPlugin(ctx, greeting.GreeterPluginOption{})
	if err != nil {
		panic(err)
	}

	pl, err := p.Load(ctx, "plugin/plugin.wasm")
	if err != nil {
		panic(err)
	}

	reply, err := pl.Greet(ctx, greeting.GreetRequest{
		Name: "test-with-go-plugin",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.GetMessage())
}
