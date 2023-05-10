package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniShopCoachListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMiniShopCoachListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniShopCoachListLogic {
	return &MiniShopCoachListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MiniShopCoachListLogic) MiniShopCoachList(req *types.CoachListRequest) (resp *types.CoachListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
