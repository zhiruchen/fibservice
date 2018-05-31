package server

import (
	"context"
	"errors"
	"reflect"
	"testing"

	pb "github.com/zhiruchen/fibservice/fibpb"
)

func TestServerGetFibNRecur(t *testing.T) {
	s := &FibServer{}

	cases := []struct {
		num int32
		ns  []int32
		err error
	}{
		{
			num: -1,
			ns:  nil,
			err: errors.New("Invalid Num"),
		},
		{
			num: 0,
			ns:  []int32{0},
			err: nil,
		},
		{
			num: 1,
			ns:  []int32{0, 1},
			err: nil,
		},
		{
			num: 2,
			ns:  []int32{0, 1, 1},
			err: nil,
		},
		{
			num: 5,
			ns:  []int32{0, 1, 1, 2, 3, 5},
			err: nil,
		},
		{
			num: 10,
			ns:  []int32{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
			err: nil,
		},
	}

	ctx := context.Background()
	for _, c := range cases {
		res, err := s.GetFibNRecur(ctx, &pb.GetFibNRecurReq{Num: c.num})

		if err != nil {
			if c.err == nil {
				t.Errorf("Expect error: %v, get nil\n", err)
			}

			if !(err.Error() == c.err.Error() && res == nil) {
				t.Errorf("Expect error: %v, get: %v\n", c.err, err)
			}
		}

		if res != nil {
			if !reflect.DeepEqual(res.FibNums, c.ns) {
				t.Errorf("Expect %v, get: %v\n", c.ns, res.FibNums)
			}
		}
	}
}

func TestServerGetFibNIter(t *testing.T) {
	s := &FibServer{}

	cases := []struct {
		num int32
		ns  []int32
		err error
	}{
		{
			num: -1,
			ns:  nil,
			err: errors.New("Invalid Num"),
		},
		{
			num: 0,
			ns:  []int32{0},
			err: nil,
		},
		{
			num: 1,
			ns:  []int32{0, 1},
			err: nil,
		},
		{
			num: 2,
			ns:  []int32{0, 1, 1},
			err: nil,
		},
		{
			num: 5,
			ns:  []int32{0, 1, 1, 2, 3, 5},
			err: nil,
		},
		{
			num: 10,
			ns:  []int32{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
			err: nil,
		},
	}

	ctx := context.Background()
	for _, c := range cases {
		res, err := s.GetFibNIter(ctx, &pb.GetFibNIterReq{Num: c.num})

		if err != nil {
			if c.err == nil {
				t.Errorf("Expect error: %v, get nil\n", err)
			}

			if !(err.Error() == c.err.Error() && res == nil) {
				t.Errorf("Expect error: %v, get: %v\n", c.err, err)
			}
		}

		if res != nil {
			if !reflect.DeepEqual(res.FibNums, c.ns) {
				t.Errorf("Expect %v, get: %v\n", c.ns, res.FibNums)
			}
		}
	}
}

func BenchmarkServerGetFibNRecur(b *testing.B) {
	s := &FibServer{}
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		s.GetFibNRecur(ctx, &pb.GetFibNRecurReq{Num: 20})
	}
}

func BenchmarkServerGetFibNIter(b *testing.B) {
	s := &FibServer{}
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		s.GetFibNIter(ctx, &pb.GetFibNIterReq{Num: 20})
	}
}
