package main

import (
	"fmt"
	"time"

	"context_demo/util"
)

func Parent() {
	defer fmt.Println("Parent exit")
	ch := make(chan struct{})
	go Child(ch)
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
}

func Child(ch chan struct{}) {
	defer fmt.Println("Child exit")
	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("child")
		}
	}
}

func main() {
	Parent()
	util.Block()
}
