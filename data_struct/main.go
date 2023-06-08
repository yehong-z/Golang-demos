package main

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	cnt := make(map[string]int)
	for _, row := range grid {
		cnt[fmt.Sprintf("%v", row)]++
	}
	res := 0
	for j := 0; j < n; j++ {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = grid[i][j]
		}
		res += cnt[fmt.Sprintf("%v", arr)]
	}
	return res
}

func main() {
	fmt.Println("hello")
}
