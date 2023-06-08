package consul_demo

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println(GetConfig())
}
