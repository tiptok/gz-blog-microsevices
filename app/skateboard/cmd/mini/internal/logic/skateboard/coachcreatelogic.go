package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoachCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoachCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoachCreateLogic {
	return &CoachCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoachCreateLogic) CoachCreate(req *types.CoachCreateRequest) (resp *types.CoachCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
