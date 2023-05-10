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

type SystemShopUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemShopUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemShopUpdateLogic {
	return &SystemShopUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemShopUpdateLogic) SystemShopUpdate(req *types.ShopUpdateRequest) (resp *types.ShopUpdateResponse, err error) {
	conn := l.svcCtx.DefaultDBConn()
	shop, err := l.svcCtx.ShopRepository.FindOne(l.ctx, conn, int64(req.Id))
	// 捕获错误
	if err == domain.ErrNotFound {
		return nil, xerr.NewErrMsg(fmt.Sprintf("未找到用户记录 id:%d", req.Id))
	}
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	shop.Name = req.Shop.Name
	shop.Introduction = req.Shop.Introduction
	shop.Rank = req.Shop.Rank
	shop.Address = types.GlobalAddressConvertor.EntityToDomain(req.Shop.Address)

	shop, err = l.svcCtx.ShopRepository.Update(l.ctx, conn, shop)
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	resp = &types.ShopUpdateResponse{}
	return
}
