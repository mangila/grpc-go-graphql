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
	service.UnimplementedProductQueryServiceServer
}

func (s *GrpcQueryService) GetProduct(context.Context, *wrapperspb.StringValue) (*model.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (s *GrpcQueryService) GetAllProducts(context.Context, *emptypb.Empty) (*service.GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
