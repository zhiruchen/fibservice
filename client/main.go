package main

import (
	"context"
	"log"
	"time"

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

	s := time.Now()
	iteratorRes, err := c.GetFibNIter(ctx, &pb.GetFibNIterReq{Num: 15})
	e := time.Now()

	log.Println("time: ", e.Sub(s).Seconds(), iteratorRes.FibNums, err)

	s = time.Now()
	recurRes, err := c.GetFibNRecur(ctx, &pb.GetFibNRecurReq{Num: 15})
	e = time.Now()
	log.Println("time: ", e.Sub(s).Seconds(), recurRes.FibNums, err)
}
