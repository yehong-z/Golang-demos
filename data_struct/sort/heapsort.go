package sort

import (
	"Datastruct/util"
)

func HeapSort(arr []int) {
	BuildHeap(arr)
	for i := 0; i < len(arr); i++ {
		util.Swap(arr, len(arr)-1-i, 0)
		Down(arr[0:len(arr)-1-i], 0)
	}
	util.Reverse(arr, 0, len(arr)-1)
}

func BuildHeap(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		Down(arr, i)
	}
}

func Down(arr []int, i int) {
	minIndex := i
	if 2*i+1 < len(arr) && arr[2*i+1] < arr[minIndex] {
		minIndex = 2*i + 1
	}
	if 2*i+2 < len(arr) && arr[2*i+2] < arr[minIndex] {
		minIndex = 2*i + 2
	}
	if i != minIndex {
		util.Swap(arr, i, minIndex)
		Down(arr, minIndex)
	}
}
