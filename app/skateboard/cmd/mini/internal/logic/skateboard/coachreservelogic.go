package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoachReserveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoachReserveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoachReserveLogic {
	return &CoachReserveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoachReserveLogic) CoachReserve(req *types.CoachReserveRequest) (resp *types.CoachReserveResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
