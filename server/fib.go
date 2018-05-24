package server

import (
	"context"
	"errors"

	pb "github.com/zhiruchen/fibservice/fibpb"
)

type FibServer struct{}

func (s *FibServer) GetFibNRecur(ctx context.Context, req *pb.GetFibNRecurReq) (*pb.GetFibNRecurRes, error) {
	if req.Num < 0 {
		return nil, errors.New("Invalid Num")
	}

	var ns []int32
	var i int32
	for i = 0; i <= req.Num; i++ {
		ns = append(ns, fibRecur(i))
	}
	return &pb.GetFibNRecurRes{FibNums: ns}, nil
}

func (s *FibServer) GetFibNIter(ctx context.Context, req *pb.GetFibNIterReq) (*pb.GetFibNIterRes, error) {
	if req.Num < 0 {
		return nil, errors.New("Invalid Num")
	}

	result := fibIter(req.Num)
	return &pb.GetFibNIterRes{FibNums: result}, nil
}

func fibRecur(n int32) int32 {
	ns := []int32{0, 1, 1}

	if n <= 2 {
		return ns[int(n)]
	}

	n1 := fibRecur(n - 1)
	n2 := fibRecur(n - 2)

	return n1 + n2
}

func fibIter(n int32) []int32 {
	ns := []int32{0, 1, 1}
	if n <= 2 {
		return ns[:n+1]
	}

	result := make([]int32, n+1)
	for i := 0; i <= 2; i++ {
		result[i] = ns[i]
	}

	var i int32
	for i = 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result
}
