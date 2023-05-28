package main

import (
	"fmt"
	"time"

	"context_demo/util"
)

func Parent() {
	defer fmt.Println("Parent exit")
	go Child()
}

func Child() {
	defer fmt.Println("Child exit")
	for {
		time.Sleep(time.Second)
		fmt.Println("child")
	}
}

func main() {
	Parent()
	util.Block()
}
