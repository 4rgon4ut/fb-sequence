package protobuff

import (
	"context"
	"log"
	"net"

	"github.com/bestpilotingalaxy/fbs-test-case/config"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/math"
	"github.com/bestpilotingalaxy/fbs-test-case/internal/redis"
	fib_proto "github.com/bestpilotingalaxy/fbs-test-case/proto"
	"google.golang.org/grpc"
)

// FibonaciServer ...
type FibonaciServer struct {
	fib_proto.UnimplementedFibonnaciServer

	Config *config.GRPCServer
}

// NewServer ...
func NewServer(c *config.GRPCServer) *grpc.Server {
	s := grpc.NewServer()
	srv := &FibonaciServer{Config: c}
	fib_proto.RegisterFibonnaciServer(s, srv)
	return s
}

// RunServer ...
func RunServer(s *grpc.Server, port string) {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// GetFibonacciSlice ...
func (s *FibonaciServer) GetFibonacciSlice(ctx context.Context, r *fib_proto.FibRequest) (*fib_proto.FibResponse, error) {
	res, err := redis.GetFromCache(r.Start, r.End)
	if err != nil {
		calculatedRes, toCache := math.FibonacciBig(r.Start, r.End)
		go redis.Client.ZAddNXMany(toCache)
		return &fib_proto.FibResponse{Result: calculatedRes}, nil
	}
	return &fib_proto.FibResponse{Result: res}, nil
}
