package convert

import (
	"fmt"
	"testing"
)

type A struct {
	x    int
	y    int
	z    int
	next *B
}

type B struct {
	X int
	Y int
}

type C struct {
	X int
	Y int
	Z int
}

func TestNameOf(t *testing.T) {
	test := A{x: 1}
	PrintNameOfField(&test)
	testB := B{
		X: 1,
		Y: 2,
	}
	testC := C{}
	err := mapStruct(&testB, &testC)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(testC)
}
