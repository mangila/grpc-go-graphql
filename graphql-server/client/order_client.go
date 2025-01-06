package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"shared/logger"
	"shared/model"
	"shared/service"
	"time"
)

type orderClient struct {
	query   service.OrderQueryServiceClient
	command service.OrderCommandServiceClient
}

var (
	orderClientInstance *orderClient
)

// InitOrderClient - init the order service grpc client
func InitOrderClient(orderServiceAddr string) {
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
	conn, err := grpc.NewClient(orderServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(connParams),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithUserAgent("graphql-server"))
	if err != nil {
		logger.Logger.Fatalf("connection failed: %v", err)
	}

	orderClientInstance = &orderClient{
		query:   service.NewOrderQueryServiceClient(conn),
		command: service.NewOrderCommandServiceClient(conn),
	}
	logger.Logger.Infof("connected to ORDER_CLIENT - %s", orderServiceAddr)
}

func GetOrder(orderId string) (*model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	return orderClientInstance.query.GetOrder(
		ctx,
		wrapperspb.String(orderId),
		grpc.WaitForReady(true),
		grpc.UseCompressor("gzip"))
}

func GetAllOrders() (*service.GetAllOrdersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	return orderClientInstance.query.GetAllOrders(
		ctx,
		&emptypb.Empty{},
		grpc.WaitForReady(true),
		grpc.UseCompressor("gzip"))
}
