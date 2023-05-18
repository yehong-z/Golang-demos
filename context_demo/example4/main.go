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

func main() {
	ParentWithTimeout()
	fmt.Println("-------------")
	ParentWithDeadline()
}
