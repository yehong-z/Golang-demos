package questions

import (
	"fmt"
	"testing"
)

func TestSequentialPrint(t *testing.T) {
	SequentialPrint(5)

}

func TestSequentialPrint2(t *testing.T) {
	SequentialPrint2(5)

}

func TestQuickSelect(t *testing.T) {
	res := quickselect2([]int{1, 5, 2, 4, 7, 6, 3}, 3)
	fmt.Println(res)
}
