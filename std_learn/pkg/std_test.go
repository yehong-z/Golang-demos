package pkg

import (
	"fmt"
	"testing"
)

func TestGetRandByte(t *testing.T) {
	var a []byte
	fmt.Println(a, len(a))
	a = make([]byte, 0)
	fmt.Println(a, len(a))
	a = nil
	fmt.Println(a, len(a))
}

func TestSlice(t *testing.T) {
	var a []int
	for _, i := range a {
		fmt.Println(i)
	}
}
