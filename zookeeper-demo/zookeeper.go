package zookeeper_demo

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func HelloZooKeeper() {
	hosts := []string{"10.20.121.247:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	path := "/test"
	data := []byte("hello world")
	flags := int32(0)

	// 创建数据节点
	_, err = conn.Create(path, data, flags, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created node: %s\n", path)

	// 读取数据节点
	bytes, _, err := conn.Get(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Node value: %s\n", bytes)

	// 删除数据节点
	err = conn.Delete(path, -1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted node: %s\n", path)
}
