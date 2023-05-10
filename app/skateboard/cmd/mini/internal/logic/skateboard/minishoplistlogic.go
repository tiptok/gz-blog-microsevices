package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniShopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMiniShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniShopListLogic {
	return &MiniShopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MiniShopListLogic) MiniShopList(req *types.ShopListRequest) (resp *types.ShopListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
