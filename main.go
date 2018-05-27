package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/zhiruchen/fibservice/fibpb"
	"github.com/zhiruchen/fibservice/server"
)

func main() {
	port := flag.String("p", ":8898", "listen port")
	flag.Parse()

	if *port == "" {
		log.Fatalln("port is empty")
	}

	lis, err := net.Listen("tcp", *port)
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
