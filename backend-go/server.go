package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 8801
// Hash 8082
// Hash 5817
// Hash 9406
// Hash 8924
// Hash 1995
// Hash 4222
// Hash 4322
// Hash 7509
// Hash 2278
// Hash 7749
// Hash 8468
// Hash 1987
// Hash 6563
// Hash 3156
// Hash 6168
// Hash 4630
// Hash 3415
// Hash 9489
// Hash 5151
// Hash 2491
// Hash 6311
// Hash 7253
// Hash 6620
// Hash 1285
// Hash 3021
// Hash 6296
// Hash 1777
// Hash 5162
// Hash 5574
// Hash 8263
// Hash 6229
// Hash 1538
// Hash 9717
// Hash 2270
// Hash 7113
// Hash 2331
// Hash 7727
// Hash 4504
// Hash 4729
// Hash 8812
// Hash 1088
// Hash 5678
// Hash 6361
// Hash 5349
// Hash 1365
// Hash 9449
// Hash 9889
// Hash 3303
// Hash 1112
// Hash 5337
// Hash 5927
// Hash 1352
// Hash 6218
// Hash 4668
// Hash 3585
// Hash 8211
// Hash 1911
// Hash 5011
// Hash 7248
// Hash 8626
// Hash 8631
// Hash 8396
// Hash 4804
// Hash 6534
// Hash 4584
// Hash 3584
// Hash 2462
// Hash 7746
// Hash 4899
// Hash 8533
// Hash 1615
// Hash 6766
// Hash 5142
// Hash 3660
// Hash 6933
// Hash 9731
// Hash 6511
// Hash 4090
// Hash 8230
// Hash 1775
// Hash 1806
// Hash 6756
// Hash 8946
// Hash 3216
// Hash 6541
// Hash 1764
// Hash 9545
// Hash 6550
// Hash 1227
// Hash 1234
// Hash 9796
// Hash 7595
// Hash 1018
// Hash 9086
// Hash 8589
// Hash 7453
// Hash 6136
// Hash 1384
// Hash 2626
// Hash 4052
// Hash 1759
// Hash 9095
// Hash 8675
// Hash 7456
// Hash 3266
// Hash 1773
// Hash 5990
// Hash 7048
// Hash 5110
// Hash 2967
// Hash 4617
// Hash 6741
// Hash 9034
// Hash 4884
// Hash 1474
// Hash 7778
// Hash 7320
// Hash 2576
// Hash 9915
// Hash 6267
// Hash 8789
// Hash 2236
// Hash 7887
// Hash 5452
// Hash 7470
// Hash 8921
// Hash 8096
// Hash 1846
// Hash 5994
// Hash 7489
// Hash 9996
// Hash 4326
// Hash 6228
// Hash 9595
// Hash 1034
// Hash 5337
// Hash 3538
// Hash 8863
// Hash 5018
// Hash 4657
// Hash 8418
// Hash 6579
// Hash 6892
// Hash 6286
// Hash 8045
// Hash 6564
// Hash 6826
// Hash 3157
// Hash 4898
// Hash 8226
// Hash 2328
// Hash 7662
// Hash 3677
// Hash 7503