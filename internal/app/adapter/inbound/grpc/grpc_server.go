package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/asfung/ticus/internal/core/ports"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server         *grpc.Server
	articleService ports.ArticleService
	log            *logrus.Logger
	port           string
}

func NewGRPCServer(articleService ports.ArticleService, log *logrus.Logger) *GRPCServer {
	server := grpc.NewServer()

	// Register the ArticleService
	articleServer := NewArticleServer(articleService, log)
	RegisterArticleServiceServer(server, articleServer)

	// Enable reflection for debugging
	reflection.Register(server)

	return &GRPCServer{
		server:         server,
		articleService: articleService,
		log:            log,
		port:           "9090", // Default gRPC port
	}
}

func (s *GRPCServer) Start(lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
				if err != nil {
					s.log.Fatalf("Failed to listen: %v", err)
				}

				s.log.Infof("gRPC server starting on port %s", s.port)
				if err := s.server.Serve(lis); err != nil {
					s.log.Fatalf("Failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.log.Info("Stopping gRPC server...")
			s.server.GracefulStop()
			return nil
		},
	})
}
