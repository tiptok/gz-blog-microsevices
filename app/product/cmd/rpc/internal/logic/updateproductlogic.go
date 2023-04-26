package logic

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/product/v1"
	"github.com/tiptok/gz-blog-microsevices/app/product/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *v1.UpdateProductRequest) (*v1.UpdateProductResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.UpdateProductResponse{}, nil
}
