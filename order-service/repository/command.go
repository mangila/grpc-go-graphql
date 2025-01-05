package repository

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"shared/model"
	"shared/service"
)

// GrpcCommandService - interface to pick implement registered rpc command calls
type GrpcCommandService struct {
	service.UnimplementedOrderCommandServiceServer
}

func (s *GrpcCommandService) CreateOrder(context.Context, *service.CreateOrderRequest) (*model.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (s *GrpcCommandService) UpdateOrder(context.Context, *service.UpdateOrderRequest) (*model.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (s *GrpcCommandService) DeleteOrder(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
