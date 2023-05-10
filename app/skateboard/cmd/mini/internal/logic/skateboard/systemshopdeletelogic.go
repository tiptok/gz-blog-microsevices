package skateboard

import (
	"context"
	"fmt"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/domain"
	"github.com/tiptok/gz-blog-microsevices/pkg/xerr"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemShopDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemShopDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemShopDeleteLogic {
	return &SystemShopDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemShopDeleteLogic) SystemShopDelete(req *types.ShopDeleteRequest) (resp *types.ShopDeleteResponse, err error) {
	conn := l.svcCtx.DefaultDBConn()
	shop, err := l.svcCtx.ShopRepository.FindOne(l.ctx, conn, int64(req.Id))
	// 捕获错误
	if err == domain.ErrNotFound {
		return nil, xerr.NewErrMsg(fmt.Sprintf("未找到用户记录 id:%d", req.Id))
	}
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	_, err = l.svcCtx.ShopRepository.Delete(l.ctx, conn, shop)
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	resp = &types.ShopDeleteResponse{}
	return
}
