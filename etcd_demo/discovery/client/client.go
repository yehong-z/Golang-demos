package client

import (
	"etcd_demo/discovery"
	"fmt"
	eclient "go.etcd.io/etcd/client/v3"
	eresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	// 标准库
	"context"
)

const MyService = "zyh1/demo"
const MyEtcdURL = "http://121.36.89.81:2379"

func Main() {
	// 创建 etcd 客户端
	etcdClient, _ := eclient.NewFromURL(MyEtcdURL)

	// 创建 etcd 实现的 grpc 服务注册发现模块 resolver
	etcdResolverBuilder, _ := eresolver.NewBuilder(etcdClient)

	// 拼接服务名称，需要固定义 etcd:/// 作为前缀
	etcdTarget := fmt.Sprintf("etcd:///%s", MyService)

	// 创建 grpc 连接代理
	conn, err := grpc.Dial(
		// 服务名称
		etcdTarget,
		// 注入 etcd resolver
		grpc.WithResolvers(etcdResolverBuilder),
		// 声明使用的负载均衡策略为 roundrobin     ,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(err.Error())
	}

	defer conn.Close()

	// 创建 grpc 客户端
	client := discovery.NewHelloClient(conn)

	for {
		// 发起 grpc 请求
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		resp, err := client.BasicRPC(ctx, &discovery.Request{
			Name: "xiaoxuxiansheng",
		})

		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("resp: %+v\n", resp)
		// 每隔 1s 发起一轮请求
		<-time.After(time.Second)
	}
}
