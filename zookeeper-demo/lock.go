package zookeeper_demo

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type ZookeeperLock struct {
	c        *zk.Conn
	lockPath string
}

func NewZookeeperClient(lockPath string) *ZookeeperLock {
	rand.Seed(time.Now().UnixNano())
	conn, _, err := zk.Connect([]string{"10.20.121.247:2181"}, time.Hour)
	if err != nil {
		panic(err)
	}
	return &ZookeeperLock{c: conn, lockPath: lockPath}
}

func (z *ZookeeperLock) InitLock() {
	// 创建锁目录
	if _, err := z.c.Create(z.lockPath, []byte{}, 0, zk.WorldACL(zk.PermAll)); err != nil && err != zk.ErrNodeExists {
		panic(err)
	}
}

func (z *ZookeeperLock) GetLock() string {
	// 创建节点并获取锁
	nodePath := z.lockPath + "/node-"
	node, err := z.c.Create(nodePath, []byte{}, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}

	waitLockComplete := make(chan struct{})

	z.waitLock(z.c, node, waitLockComplete)

	<-waitLockComplete

	fmt.Println("get lock")
	return node
}

func (z *ZookeeperLock) ReleaseLock(node string) error {
	return z.c.Delete(node, -1)
}

func (z *ZookeeperLock) waitLock(conn *zk.Conn, node string, waitLockComplete chan struct{}) {
	for {
		children, _, ch, err := conn.ChildrenW(z.lockPath)
		if err != nil {
			panic(err)
		}

		var minNode string
		for _, child := range children {
			if strings.HasPrefix(child, "node-") {
				if minNode == "" {
					minNode = child
				} else if child < minNode {
					minNode = child
				}
			}
		}

		if minNode == node[len(z.lockPath)+1:] {
			close(waitLockComplete)
			break
		}

		select {
		case <-ch:
		case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		}
	}
}
