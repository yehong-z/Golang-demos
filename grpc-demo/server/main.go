package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/protoc"
	"log"
	"net"
)

type HelloServer struct {
	pb.UnimplementedHelloServer
}

func (s *HelloServer) BasicRPC(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("receive : %v", in.Name)
	res := in.GetName()
	for _, name := range res {
		fmt.Println(name)
	}
	return &pb.Reply{}, nil
}

func (s *HelloServer) Many(ctx context.Context, in *pb.Requests) (*pb.BatchReply, error) {
	res := in.Request
	for _, r := range res {
		fmt.Println(r.GetName())
	}
	return &pb.BatchReply{}, nil
}

//func (s *HelloServer) StreamInput(stream pb.Hello_StreamInputServer) (err error) {
//	for {
//		str, err := stream.Recv()
//		if err == io.EOF {
//			return stream.SendAndClose(&pb.Reply{Msg: str.GetName()})
//		} else {
//			log.Printf("%v", str.GetName())
//		}
//	}
//}
//
//func (s *HelloServer) StreamOutput(req *pb.Request, stream pb.Hello_StreamOutputServer) error {
//	for i := 0; i < 10; i++ {
//		err := stream.Send(&pb.Reply{Msg: req.GetName()})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func main() {
	lis, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	log.Printf("server listening at %v", lis.Addr())
	pb.RegisterHelloServer(s, &HelloServer{})
	err = s.Serve(lis)
	if err != nil {
		return
	}
}
