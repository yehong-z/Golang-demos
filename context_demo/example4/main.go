package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func Child(ctx context.Context) error {
	defer fmt.Println("child exit")

	select {
	case <-ctx.Done():
		return errors.New("timeout")
	case <-time.After(time.Second * 2):
		return nil
	}
}

func ParentWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	fmt.Println(fmt.Errorf("%w", Child(ctx)))
}

func ParentWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()
	fmt.Println(fmt.Errorf("%w", Child(ctx)))
}

func Request() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ch := make(chan error)
	go worker(ch)

	select {
	case <-ch:
		fmt.Println("ok")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

func worker(ch chan error) {
	time.Sleep(time.Second)
	ch <- nil
}

func main() {
	Request()
	// ParentWithTimeout()
	// fmt.Println("-------------")
	// ParentWithDeadline()
}
