package server

import (
	"net"

	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "abtest/abtest_proto"
)

var (
	// GRPCAddress is the address
	GRPCAddress string = ":9527"
)

// ABTestGRPCServer is a grpc server
type ABTestGRPCServer struct{}

// MustStartGRPCServer is to start the grpc server
func MustStartGRPCServer() {
	err := run()
	if err != nil {
		glog.Fatalf("run: %v", err)
	}
}

func run() error {
	// New gRPC server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			otelgrpc.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpc_validator.StreamServerInterceptor(),
		)),
	)

	reflection.Register(s)
	pb.RegisterABTestServiceServer(s, &ABTestGRPCServer{})

	lis, err := net.Listen("tcp", GRPCAddress)
	if err != nil {
		return errors.Wrap(err, "tcp listen")
	}
	err = s.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "serve gRPC server")
	}
	return nil
}
