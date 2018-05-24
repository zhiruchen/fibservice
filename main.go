package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/zhiruchen/fibservice/fibpb"
	"github.com/zhiruchen/fibservice/server"
)

func main() {
	lis, err := net.Listen("tcp", ":8898")
	if err != nil {
		log.Printf("listen error: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterFibgServer(s, &server.FibServer{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Printf("serve error: %v\n", err)
		return
	}
}
