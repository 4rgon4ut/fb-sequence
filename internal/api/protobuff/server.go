package protobuff

import (
	"context"
	"log"
	"net"

	"github.com/bestpilotingalaxy/fbs-test-case/internal/math"
	fib_proto "github.com/bestpilotingalaxy/fbs-test-case/proto"
	"google.golang.org/grpc"
)

// FibonaciServer ...
type FibonaciServer struct {
	fib_proto.UnimplementedFibonnaciServer
}

// NewServer ...
func NewServer() *grpc.Server {
	s := grpc.NewServer()
	srv := &FibonaciServer{}
	fib_proto.RegisterFibonnaciServer(s, srv)
	return s
}

// RunServer ...
func RunServer(s *grpc.Server) {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// GetFibonacciSlice ...
func (s *FibonaciServer) GetFibonacciSlice(ctx context.Context, r *fib_proto.FibRequest) (*fib_proto.FibResponse, error) {
	result := math.FibonacciBig(r.Start, r.End)
	return &fib_proto.FibResponse{Result: result}, nil
}
