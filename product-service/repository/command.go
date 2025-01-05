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
	service.UnimplementedProductCommandServiceServer
}

func (s *GrpcCommandService) CreateProduct(context.Context, *service.CommandProductRequest) (*model.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (s *GrpcCommandService) UpdateProduct(context.Context, *service.CommandProductRequest) (*model.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (s *GrpcCommandService) DeleteProduct(context.Context, *wrappers.StringValue) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
