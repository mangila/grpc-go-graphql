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
	service.UnimplementedCustomerQueryServiceServer
}

func (s *GrpcQueryService) GetCustomer(context.Context, *wrappers.StringValue) (*model.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (s *GrpcQueryService) GetAllCustomers(context.Context, *empty.Empty) (*service.GetAllCustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCustomers not implemented")
}
