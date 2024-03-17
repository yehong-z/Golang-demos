package config

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func Main() {
	// 创建 etcd 客户端连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://121.36.89.81:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()
	// 从 etcd 中获取配置信息
	resp, err := cli.Get(ctx, "a")
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range resp.Kvs {
		log.Printf("%s : %s\n", kv.Key, kv.Value)
	}

	ch := cli.Watch(ctx, "a")
	for {
		select {
		case t := <-ch:
			fmt.Println(t.Events[0].Kv.String())

		default:
			fmt.Println("default")
			time.Sleep(time.Second)
		}
	}
}
