package repository

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"shared/model"
	"shared/service"
)

// GrpcQueryService - interface to pick implement registered rpc query calls
type GrpcQueryService struct {
	service.UnimplementedOrderQueryServiceServer
}

func (s *GrpcQueryService) GetOrder(context.Context, *wrapperspb.StringValue) (*model.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}

func (s *GrpcQueryService) GetAllOrders(context.Context, *emptypb.Empty) (*service.GetAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
