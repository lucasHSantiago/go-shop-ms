package server

import (
	"github.com/lucasHSantiago/go-shop-ms/category/category"
	"github.com/lucasHSantiago/go-shop-ms/category/internal/grpc/pb"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryServer(s category.UseCase) *CategoryServer {
	return &CategoryServer{}
}
