package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopListLogic {
	return &ShopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopListLogic) ShopList(req *types.ShopListRequest) (resp *types.ShopListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
