package logic

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/product/cmd/rpc/productservice"
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/product/v1"
	"github.com/tiptok/gz-blog-microsevices/app/product/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *v1.CreateProductRequest) (*v1.CreateProductResponse, error) {
	product := &domain.Products{
		Name:        in.GetProduct().Name,
		Description: in.GetProduct().Description,
		Price:       in.GetProduct().Price,
		Unit:        in.GetProduct().Unit,
		Stock:       in.GetProduct().Stock,
	}

	var err error
	if err = transaction.UseTrans(l.ctx, l.svcCtx.DB, func(ctx context.Context, conn transaction.Conn) error {
		product, err = l.svcCtx.ProductsRepository.Insert(l.ctx, conn, product)
		return err
	}, true); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}
	return &v1.CreateProductResponse{
		Product: productservice.ProductEntityToProtobuf(product),
	}, nil
}
