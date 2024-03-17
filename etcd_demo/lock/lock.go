package lock

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Main() {
	// 创建 etcd 客户端连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://121.36.89.81:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建一个租约
	resp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		log.Fatal(err)
	}
	leaseID := resp.ID

	// 创建一个带租约的锁路径
	lockPath := "/mylock"
	for {
		// 尝试获取锁
		resp, err := cli.Txn(context.Background()).
			If(clientv3.Compare(clientv3.CreateRevision(lockPath), "=", 0)).
			Then(clientv3.OpPut(lockPath, "locked", clientv3.WithLease(leaseID))).
			Commit()
		if err != nil {
			log.Fatal(err)
		}

		if resp.Succeeded {
			log.Println("获取锁成功")
			break
		} else {
			log.Println("等待锁释放")
			time.Sleep(1 * time.Second)
		}
	}

	// 模拟持有锁期间的操作
	log.Println("执行任务...")
	time.Sleep(5 * time.Second)

	// 释放锁
	_, err = cli.Delete(context.Background(), lockPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("释放锁成功")
}
