package main

import (
	"context"
	"fmt"
)

func Child(ctx context.Context) {
	fmt.Println(ctx.Value("k1"))
}

func Parent() {
	ctx := context.WithValue(context.Background(), "k1", "v1")
	Child(ctx)
}

func main() {
	Parent()
}
