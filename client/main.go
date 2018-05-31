package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/zhiruchen/fibservice/fibpb"
)

func main() {
	address := "localhost:8898"
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFibgClient(conn)

	for i := 1; i <= 10000; i++ {
		iteratorRes, err := c.GetFibNIter(ctx, &pb.GetFibNIterReq{Num: 25})
		log.Println(iteratorRes.FibNums, err)
		recurRes, err := c.GetFibNRecur(ctx, &pb.GetFibNRecurReq{Num: 25})
		log.Println(recurRes.FibNums, err)
	}
}
