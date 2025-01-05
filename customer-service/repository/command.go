package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shared/model"
	"shared/service"
)

// GrpcCommandService - interface to pick implement registered rpc command calls
type GrpcCommandService struct {
	service.UnimplementedCustomerCommandServiceServer
}

func (s *GrpcCommandService) CreateCustomer(context.Context, *service.CommandCustomerRequest) (*model.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (s *GrpcCommandService) UpdateCustomer(context.Context, *service.CommandCustomerRequest) (*model.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (s *GrpcCommandService) DeleteCustomer(context.Context, *wrappers.StringValue) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
