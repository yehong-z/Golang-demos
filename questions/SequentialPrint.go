package questions

import (
	"context"
	"fmt"
	"sync"
)

func SequentialPrint(n int) {
	chanArray := make([]chan int, n)
	for i := range chanArray {
		chanArray[i] = make(chan int, 1)
	}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				j := <-chanArray[i]
				if j < 100 {
					fmt.Println("goroutine", i, j)
					chanArray[(i+1)%n] <- j + 1
				} else {
					chanArray[(i+1)%n] <- j + 1
					return
				}
			}
		}(i)
	}
	chanArray[0] <- 1
	wg.Wait()
}

func SequentialPrint2(n int) {
	chArray := make([]chan int, n)
	for i := range chArray {
		chArray[i] = make(chan int, 1)
	}

	wg := sync.WaitGroup{}
	wg.Add(n)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < n; i++ {
		go func(ctx context.Context, i int) {
			defer wg.Done()
			for {
				select {
				case j := <-chArray[i]:
					fmt.Println("goroutine", i+1, j)
					if j == 100 {
						cancel()
						return
					}
					chArray[(i+1)%n] <- j + 1
				case <-ctx.Done():
					return
				}
			}
		}(ctx, i)
	}
	chArray[0] <- 1
	wg.Wait()
}
