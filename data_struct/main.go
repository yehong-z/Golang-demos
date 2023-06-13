package main

import (
	"fmt"
	"sort"
)

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	l := 0
	r := int(1e9 + 10)
	var mid int
	for l < r {
		mid = (l + r) / 2
		a := price[0]
		cnt := 0
		for i := 1; i < len(price); i++ {
			if price[i]-a >= mid {
				cnt++
				a = price[i]
			}
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println("hello")
}
