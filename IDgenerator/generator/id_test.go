package generator

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	fmt.Println(getUUID())
	fmt.Println(getSnowFlakeID())
}
