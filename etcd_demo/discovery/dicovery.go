package discovery

import (
	"flag"
	"fmt"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"time"

	// 标准库
	"context"
	"net"
	// grpc
	client "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

const (
	// grpc 服务名
	MyService = "zyh/demo"
	// etcd 端口
	MyEtcdURL = "http://121.36.89.81:2379"
)

type Server struct {
	UnimplementedHelloServer
}

func Main() {
	// 接收命令行指定的 grpc 服务端口
	var port int
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
	addr := fmt.Sprintf(":%d", port)

	// 创建 tcp 端口监听器
	listener, _ := net.Listen("tcp", addr)

	// 创建 grpc server
	server := grpc.NewServer()
	RegisterHelloServer(server, &Server{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 注册 grpc 服务节点到 etcd 中
	go registerEndPointToEtcd(ctx, addr)
	// 启动 grpc 服务
	if err := server.Serve(listener); err != nil {
		fmt.Println(err)
	}
}

func registerEndPointToEtcd(ctx context.Context, addr string) {
	// 创建 etcd 客户端
	etcdClient, err := client.NewFromURL(MyEtcdURL)
	if err != nil {
		log.Println(err.Error())
	}
	etcdManager, err := endpoints.NewManager(etcdClient, MyService)
	if err != nil {
		log.Println(err.Error())
	}
	// 创建一个租约，每隔 10s 需要向 etcd 汇报一次心跳，证明当前节点仍然存活
	var ttl int64 = 10
	lease, err := etcdClient.Grant(ctx, ttl)
	if err != nil {
		log.Println(err.Error())
	}
	// 添加注册节点到 etcd 中，并且携带上租约 id
	err = etcdManager.AddEndpoint(ctx, fmt.Sprintf("%s/%s", MyService, addr), endpoints.Endpoint{Addr: addr}, client.WithLease(lease.ID))
	if err != nil {
		log.Println(err.Error())
	}
	// 每隔 5 s进行一次延续租约的动作
	for {
		select {
		case <-time.After(5 * time.Second):
			// 续约操作
			resp, err := etcdClient.KeepAliveOnce(ctx, lease.ID)
			if err != nil {
				log.Println(err.Error())
			}
			fmt.Printf("keep alive resp: %+v\n", resp)
		case <-ctx.Done():
			return
		}
	}
}
