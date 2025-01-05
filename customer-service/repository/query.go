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
	service.UnimplementedCustomerQueryServiceServer
}

func (s *GrpcQueryService) GetCustomer(context.Context, *wrapperspb.StringValue) (*model.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (s *GrpcQueryService) GetAllCustomers(context.Context, *emptypb.Empty) (*service.GetAllCustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCustomers not implemented")
}
