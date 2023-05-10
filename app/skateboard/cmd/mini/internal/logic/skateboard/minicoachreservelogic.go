package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniCoachReserveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMiniCoachReserveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniCoachReserveLogic {
	return &MiniCoachReserveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MiniCoachReserveLogic) MiniCoachReserve(req *types.CoachReserveRequest) (resp *types.CoachReserveResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
