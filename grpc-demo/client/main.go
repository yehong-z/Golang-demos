package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/protoc"
)

func BasicRCP(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.BasicRPC(ctx, &pb.Request{Name: "req"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
}

func Many(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var req []*pb.Request
	req = append(req, &pb.Request{Name: "aaa"})
	req = append(req, &pb.Request{Name: "aaa"})
	req = append(req, &pb.Request{Name: "aaa"})
	req = append(req, &pb.Request{Name: "aaa"})
	_, err := c.Many(ctx, &pb.Requests{Request: req})
	if err != nil {
		return
	}
}

func main() {
	conn, err := grpc.Dial("localhost:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	// BasicRCP(c)
	Many(c)
	// StreamInput(c)
	// StreamOutput(c)
}

//func StreamInput(c pb.HelloClient) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	stream, err := c.StreamInput(ctx)
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	for i := 0; i < 10; i++ {
//		err = stream.Send(&pb.Request{Name: "aaa"})
//		if err != nil {
//			return
//		}
//	}
//	rep, err := stream.CloseAndRecv()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//	log.Printf(rep.GetMsg())
//}
//
//func StreamOutput(c pb.HelloClient) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	stream, err := c.StreamOutput(ctx, &pb.Request{Name: "abc"})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	for {
//		rep, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatalf("client.ListFeatures failed: %v", err)
//		}
//		log.Printf("%v", rep.GetMsg())
//	}
//	log.Printf("end")
//}
