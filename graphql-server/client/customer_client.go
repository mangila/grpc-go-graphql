package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"shared/logger"
	"shared/service"
	"time"
)

type customerClient struct {
	query   service.CustomerQueryServiceClient
	command service.CustomerCommandServiceClient
}

var (
	customerClientInstance *customerClient
)

// InitCustomerClient - init the customer service grpc client
func InitCustomerClient(customerServiceAddr string) {
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

	customerClientInstance = &customerClient{
		query:   service.NewCustomerQueryServiceClient(conn),
		command: service.NewCustomerCommandServiceClient(conn),
	}
	logger.Logger.Infof("connected to CUSTOMER_SERVICE - %s", customerServiceAddr)
}
