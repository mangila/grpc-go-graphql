package server

import (
	"context"
	"customer-service/repository"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"shared/logger"
	"shared/service"
	"time"
)

// Init - init the grpc server
func Init(port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Logger.Fatalf("failed to create listener on port:  %s - %v", port, err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryLoggingInterceptor),
	)
	service.RegisterCustomerQueryServiceServer(grpcServer, &repository.GrpcQueryService{})
	service.RegisterCustomerCommandServiceServer(grpcServer, &repository.GrpcCommandService{})
	logger.Logger.Infof("customer-service running on port: %v", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Logger.Fatalf("failed to start GRPC server on port: %s - %v", port, err)
	}
}

// UnaryLoggingInterceptor - simple logging interceptor for Unary requests
func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Log request details
	entry := logger.Logger.WithFields(logrus.Fields{
		"method": info.FullMethod,
		"start":  start.Format(time.RFC3339),
	})

	// Extract metadata
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		entry = entry.WithField("metadata", md)
	}

	// Process the request
	resp, err := handler(ctx, req)

	// Log response details and duration
	duration := time.Since(start)
	if err != nil {
		entry.WithFields(logrus.Fields{
			"error":    err,
			"duration": duration.Milliseconds(),
		}).Error("Unary request failed")
	} else {
		entry.WithFields(logrus.Fields{
			"response": resp,
			"duration": duration,
		}).Info("Unary request succeeded")
	}

	return resp, err
}
