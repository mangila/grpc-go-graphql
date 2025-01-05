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
	service.UnimplementedCustomerCommandServiceServer
}

func (s *GrpcCommandService) CreateCustomer(context.Context, *service.CommandCustomerRequest) (*model.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (s *GrpcCommandService) UpdateCustomer(context.Context, *service.CommandCustomerRequest) (*model.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (s *GrpcCommandService) DeleteCustomer(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
