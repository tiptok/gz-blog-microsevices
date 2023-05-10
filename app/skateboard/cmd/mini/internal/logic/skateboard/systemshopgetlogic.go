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

type SystemShopGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemShopGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemShopGetLogic {
	return &SystemShopGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemShopGetLogic) SystemShopGet(req *types.ShopGetRequest) (resp *types.ShopGetResponse, err error) {
	conn := l.svcCtx.DefaultDBConn()
	shop, err := l.svcCtx.ShopRepository.FindOne(l.ctx, conn, int64(req.Id))
	// 捕获错误
	if err == domain.ErrNotFound {
		return nil, xerr.NewErrMsg(fmt.Sprintf("未找到用户记录 id:%d", req.Id))
	}
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	resp = &types.ShopGetResponse{
		Shop: types.GlobalShopConvertor.DomainToEntity(shop),
	}
	return
}