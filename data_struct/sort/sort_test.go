package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

const LEN = 1e7

func TestSort(t *testing.T) {
	targetArray := make([]int, LEN)
	for i := 0; i < LEN; i++ {
		targetArray[i] = rand.Int()
	}
	testArray := make([]int, LEN)
	copy(testArray, targetArray)
	sort.Ints(targetArray)

	cases := []func([]int){
		sort.Ints,
		// InsertSort,
		QuickSort,
		HeapSort,
	}
	for _, c := range cases {
		sortTypeName := runtime.FuncForPC(reflect.ValueOf(c).Pointer()).Name()
		t.Run(sortTypeName, func(t *testing.T) {
			Array := make([]int, LEN)
			copy(Array, testArray)
			c(Array)
			for i := 0; i < LEN; i++ {
				if Array[i] != targetArray[i] {
					t.Fatalf("sort error in index %v. \n", i)
				}
			}
		})
	}
}

func TestHeap(t *testing.T) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = 10 - i
	}
	for _, i := range arr {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	HeapSort(arr)
	for _, i := range arr {
		fmt.Printf("%v ", i)
	}
}
