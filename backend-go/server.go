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