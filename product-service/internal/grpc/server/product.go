package server

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-shop-ms/product/internal/grpc/pb"
	"github.com/lucasHSantiago/go-shop-ms/product/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	useCase product.UseCase
}

func NewProductServer(s product.UseCase) *ProductServer {
	return &ProductServer{
		useCase: s,
	}
}

func (s *ProductServer) Create(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	nn, err := toNewProduct(req.GetNewProduct())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid new product data: %s", err)
	}

	products, err := s.useCase.Create(ctx, nn)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %s", err)
	}

	return &pb.CreateProductResponse{Product: toProtoProduct(products)}, nil
}

func (s *ProductServer) GetAll(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	filter, err := toProductFilter(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid filter: %s", err)
	}

	products, err := s.useCase.Get(ctx, *filter, int(req.GetPageNumber()), int(req.GetPageRows()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get products: %s", err)
	}

	return &pb.GetAllProductsResponse{Products: toProtoProduct(products)}, nil
}

// -------------------------------------------------------------------------
// Helper functions

func toNewProduct(newProducts []*pb.NewProduct) ([]product.NewProduct, error) {
	nn := make([]product.NewProduct, 0, len(newProducts))
	for _, np := range newProducts {
		categoryId, err := uuid.Parse(np.GetCategoryId())
		if err != nil {
			return nil, fmt.Errorf("category_id is not a UUID: %w", err)
		}

		newProd := product.NewProduct{
			Name:        np.GetName(),
			Description: np.GetDescription(),
			Price:       np.GetPrice(),
			CategoryId:  categoryId,
		}
		nn = append(nn, newProd)
	}
	return nn, nil
}

func toProductFilter(req *pb.GetAllProductsRequest) (*product.Filter, error) {
	filter := &product.Filter{
		Name:  req.Name,
		Price: req.Price,
	}

	if req.CategoryId != nil {
		var err error
		*filter.CategoryId, err = uuid.Parse(req.GetCategoryId())
		if err != nil {
			return nil, fmt.Errorf("category_id is not a UUID: %w", err)
		}
	}

	return filter, nil
}

func toProtoProduct(products []*product.Product) []*pb.Product {
	protoProducts := make([]*pb.Product, 0, len(products))
	for _, prd := range products {
		protoProducts = append(protoProducts, &pb.Product{
			Id:          prd.ID.String(),
			Name:        prd.Name,
			Description: prd.Description,
			Price:       prd.Price,
			CategoryId:  prd.CategoryId.String(),
			CreatedAt:   prd.Created_at.Format(time.RFC3339),
		})
	}
	return protoProducts
}
