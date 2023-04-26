package logic

import (
	"context"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/product/cmd/rpc/productservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/product/v1"
	"github.com/tiptok/gz-blog-microsevices/app/product/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *v1.GetProductRequest) (*v1.GetProductResponse, error) {
	conn := l.svcCtx.DefaultDBConn()
	product, err := l.svcCtx.ProductsRepository.FindOne(l.ctx, conn, in.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("not found product id %v", in.GetId()))
	}
	return &v1.GetProductResponse{
		Product: productservice.ProductEntityToProtobuf(product),
	}, nil
}
