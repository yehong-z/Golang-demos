package sort

import (
	"Datastruct/util"
	_ "Datastruct/util"
)

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			util.Swap(arr, j, j-1)
			j--
		}
	}
}
