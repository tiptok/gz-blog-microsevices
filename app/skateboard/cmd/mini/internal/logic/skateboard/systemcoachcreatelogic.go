package skateboard

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemCoachCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemCoachCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemCoachCreateLogic {
	return &SystemCoachCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemCoachCreateLogic) SystemCoachCreate(req *types.CoachCreateRequest) (resp *types.CoachCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
