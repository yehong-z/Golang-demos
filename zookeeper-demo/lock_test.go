package zookeeper_demo

import (
	"fmt"
	"testing"
)

func TestLock(t *testing.T) {
	z := NewZookeeperClient("/lock")
	defer z.c.Close()
	z.InitLock()
	node := z.GetLock()
	err := z.ReleaseLock(node)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("release lock")
}
