package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
	"shared/logger"
	"shared/service"
	"time"
)

type grpcClient struct {
	query   service.CustomerQueryServiceClient
	command service.CustomerCommandServiceClient
}

var (
	client *grpcClient
)

// InitCustomerService - init the customer service grpc client
func InitCustomerService(customerServiceAddr string) {
	connParams := grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  1 * time.Second,  // Initial delay
			MaxDelay:   10 * time.Second, // Max delay for retries
			Multiplier: 1.6,              // Multiplier for exponential backoff
		},
	}
	keepaliveParams := keepalive.ClientParameters{
		Time:                10 * time.Second, // Ping server every 10 seconds if no activity
		Timeout:             5 * time.Second,  // Wait 5 seconds for ping ack before assuming connection is dead
		PermitWithoutStream: true,             // Send keepalive pings even without active RPCs
	}
	conn, err := grpc.NewClient(customerServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(connParams),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithUserAgent("graphql-server"))
	if err != nil {
		logger.Logger.Fatalf("connection failed: %v", err)
	}

	client = &grpcClient{
		query:   service.NewCustomerQueryServiceClient(conn),
		command: service.NewCustomerCommandServiceClient(conn),
	}
	logger.Logger.Infof("connected to CUSTOMER_SERVICE - %s", customerServiceAddr)
}

func GetAllCustomers() (*service.GetAllCustomersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	return client.query.GetAllCustomers(ctx, &emptypb.Empty{})
}
