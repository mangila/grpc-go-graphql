package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shared/model"
	"shared/service"
)

// GrpcQueryService - interface to pick implement registered rpc query calls
type GrpcQueryService struct {
	service.UnimplementedOrderQueryServiceServer
}

func (s *GrpcQueryService) GetOrder(context.Context, *wrappers.StringValue) (*model.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}

func (s *GrpcQueryService) GetAllOrders(context.Context, *empty.Empty) (*service.GetAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
