package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/product/v1"
	"github.com/tiptok/gz-blog-microsevices/app/product/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProductsLogic) ListProducts(in *v1.ListProductsRequest) (*v1.ListProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.ListProductsResponse{}, nil
}
