# cpu profile
通过 cpu profile 可以定位耗时
```shell
go tool pprof http://localhost:8899/internal/debug/pprof/profile
Fetching profile over HTTP from http://localhost:8899/internal/debug/pprof/profile
Local symbolization failed for fibservice: open /data/apps/fibservice/bin/fibservice: no such file or directory
Some binary filenames not available. Symbolization may be incomplete.
Try setting PPROF_BINARY_PATH to the search path for local binaries.
Saved profile in /Users/zhiruchen/pprof/pprof.fibservice.samples.cpu.004.pb.gz
File: fibservice
Type: cpu
Time: May 31, 2018 at 9:25pm (CST)
Duration: 30s, Total samples = 2.07s ( 6.90%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.94s, 93.72% of 2.07s total
Dropped 52 nodes (cum <= 0.01s)
Showing top 10 nodes out of 74
      flat  flat%   sum%        cum   cum%
     1.29s 62.32% 62.32%      1.29s 62.32%  github.com/zhiruchen/fibservice/server.fibRecur
     0.17s  8.21% 70.53%      0.17s  8.21%  runtime._ExternalCode
     0.16s  7.73% 78.26%      0.16s  7.73%  runtime.usleep
     0.14s  6.76% 85.02%      0.14s  6.76%  runtime.futex
     0.10s  4.83% 89.86%      0.14s  6.76%  syscall.Syscall
     0.02s  0.97% 90.82%      0.02s  0.97%  runtime.epollwait
     0.02s  0.97% 91.79%      0.02s  0.97%  runtime.pcvalue
     0.02s  0.97% 92.75%      0.20s  9.66%  runtime.sysmon
     0.01s  0.48% 93.24%      0.07s  3.38%  github.com/zhiruchen/fibservice/vendor/golang.org/x/net/http2.(*Framer).ReadFrame
     0.01s  0.48% 93.72%      0.03s  1.45%  github.com/zhiruchen/fibservice/vendor/google.golang.org/grpc.(*Server).sendResponse
(pprof) top10 -cum
Showing nodes accounting for 1.48s, 71.50% of 2.07s total
Dropped 52 nodes (cum <= 0.01s)
Showing top 10 nodes out of 74
      flat  flat%   sum%        cum   cum%
         0     0%     0%      1.33s 64.25%  github.com/zhiruchen/fibservice/vendor/google.golang.org/grpc.(*Server).serveStreams.func1.1
         0     0%     0%      1.32s 63.77%  github.com/zhiruchen/fibservice/vendor/google.golang.org/grpc.(*Server).handleStream
         0     0%     0%      1.32s 63.77%  github.com/zhiruchen/fibservice/vendor/google.golang.org/grpc.(*Server).processUnaryRPC
         0     0%     0%      1.29s 62.32%  github.com/zhiruchen/fibservice/fibpb._Fibg_GetFibNRecur_Handler
         0     0%     0%      1.29s 62.32%  github.com/zhiruchen/fibservice/server.(*FibServer).GetFibNRecur
     1.29s 62.32% 62.32%      1.29s 62.32%  github.com/zhiruchen/fibservice/server.fibRecur
         0     0% 62.32%      0.20s  9.66%  runtime.mstart
         0     0% 62.32%      0.20s  9.66%  runtime.mstart1
     0.02s  0.97% 63.29%      0.20s  9.66%  runtime.sysmon
     0.17s  8.21% 71.50%      0.17s  8.21%  runtime._ExternalCode
(pprof) list GetFibNRecur
Total: 2.07s
ROUTINE ======================== github.com/zhiruchen/fibservice/fibpb._Fibg_GetFibNRecur_Handler in /Users/zhiruchen/go/src/github.com/zhiruchen/fibservice/fibpb/fibs.pb.go
         0      1.29s (flat, cum) 62.32% of Total
         .          .    169:	in := new(GetFibNRecurReq)
         .          .    170:	if err := dec(in); err != nil {
         .          .    171:		return nil, err
         .          .    172:	}
         .          .    173:	if interceptor == nil {
         .      1.29s    174:		return srv.(FibgServer).GetFibNRecur(ctx, in)
         .          .    175:	}
         .          .    176:	info := &grpc.UnaryServerInfo{
         .          .    177:		Server:     srv,
         .          .    178:		FullMethod: "/fibpb.fibg/GetFibNRecur",
         .          .    179:	}
ROUTINE ======================== github.com/zhiruchen/fibservice/server.(*FibServer).GetFibNRecur in /Users/zhiruchen/go/src/github.com/zhiruchen/fibservice/server/fib.go
         0      1.29s (flat, cum) 62.32% of Total
         .          .     15:	}
         .          .     16:
         .          .     17:	var ns []int32
         .          .     18:	var i int32
         .          .     19:	for i = 0; i <= req.Num; i++ {
         .      1.29s     20:		ns = append(ns, fibRecur(i))
         .          .     21:	}
         .          .     22:	return &pb.GetFibNRecurRes{FibNums: ns}, nil
         .          .     23:}
         .          .     24:
         .          .     25:func (s *FibServer) GetFibNIter(ctx context.Context, req *pb.GetFibNIterReq) (*pb.GetFibNIterRes, error) {
(pprof) list fibRecur
Total: 2.07s
ROUTINE ======================== github.com/zhiruchen/fibservice/server.fibRecur in /Users/zhiruchen/go/src/github.com/zhiruchen/fibservice/server/fib.go
     1.29s      3.45s (flat, cum) 166.67% of Total
         .          .     29:
         .          .     30:	result := fibIter(req.Num)
         .          .     31:	return &pb.GetFibNIterRes{FibNums: result}, nil
         .          .     32:}
         .          .     33:
     300ms      300ms     34:func fibRecur(n int32) int32 {
     140ms      140ms     35:	ns := []int32{0, 1, 1}
         .          .     36:
      70ms       70ms     37:	if n <= 2 {
     100ms      100ms     38:		return ns[int(n)]
         .          .     39:	}
         .          .     40:
     380ms      1.29s     41:	n1 := fibRecur(n - 1)
      40ms      1.29s     42:	n2 := fibRecur(n - 2)
         .          .     43:
     260ms      260ms     44:	return n1 + n2
         .          .     45:}
         .          .     46:
         .          .     47:func fibIter(n int32) []int32 {
         .          .     48:	ns := []int32{0, 1, 1}
         .          .     49:	if n <= 2 {
(pprof) 

```

从cpu profile(` top10 -cum`)中看到除了grpc server的几个方法，server.(*FibServer).GetFibNRecur是耗时最多的一个handler。

`list GetFibNRecur` 通过这个命令定位到cpu主要消耗在`srv.(FibgServer).GetFibNRecur(ctx, in)`方法上。
`list fibRecur` 定位是递归调用自己耗时最多。递归自然是很慢。

## heap profile

```shell
go test -run=^$ -bench=. -memprofile=mem0.out
```

```shell
go tool pprof --alloc_space server.test mem0.out
File: server.test
Type: alloc_space
Time: May 31, 2018 at 10:00pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1477.10MB, 100% of 1477.10MB total
      flat  flat%   sum%        cum   cum%
 1146.59MB 77.62% 77.62%  1146.59MB 77.62%  github.com/zhiruchen/fibservice/server.fibIter
  330.51MB 22.38%   100%  1477.10MB   100%  github.com/zhiruchen/fibservice/server.(*FibServer).GetFibNIter
         0     0%   100%  1477.10MB   100%  github.com/zhiruchen/fibservice/server.BenchmarkServerGetFibNIter
         0     0%   100%  1477.10MB   100%  testing.(*B).launch
         0     0%   100%  1477.10MB   100%  testing.(*B).runN
```

从top的结果来看server.fibIter分配的内存最多

```shell
list server.fibIter
Total: 1.44GB
ROUTINE ======================== github.com/zhiruchen/fibservice/server.fibIter in /Users/zhiruchen/go/src/github.com/zhiruchen/fibservice/server/fib.go
    1.12GB     1.12GB (flat, cum) 77.62% of Total
         .          .     43:
         .          .     44:	return n1 + n2
         .          .     45:}
         .          .     46:
         .          .     47:func fibIter(n int32) []int32 {
     177MB      177MB     48:	ns := []int32{0, 1, 1}
         .          .     49:	if n <= 2 {
         .          .     50:		return ns[:n+1]
         .          .     51:	}
         .          .     52:
  969.59MB   969.59MB     53:	result := make([]int32, n+1)
         .          .     54:	for i := 0; i <= 2; i++ {
         .          .     55:		result[i] = ns[i]
         .          .     56:	}
         .          .     57:
         .          .     58:	var i int32

```
从list 列出的代码中发现`result := make([]int32, n+1)`累积分配了最多的内存。可以优化这行代码