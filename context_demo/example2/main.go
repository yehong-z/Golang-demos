package main

import (
	"context"
	"fmt"
	"time"

	"context_demo/util"
)

func Parent() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		fmt.Println("Parent exit")
		cancel()
	}()
	go Child(ctx)
}

func Child(ctx context.Context) {
	defer fmt.Println("Child exit")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	Parent()
	util.Block()
}
