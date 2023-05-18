package main

import (
	"context_demo/util"
	"fmt"
	"time"
)

func Parent() {
	defer fmt.Println("Parent exit")
	go Child()
}

func Child() {
	defer fmt.Println("Parent exit")
	for {
		time.Sleep(time.Second)
		fmt.Println("child")
	}
}

func main() {
	Parent()
	util.Block()
}
