package skateboard

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/transaction"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemShopSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemShopSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemShopSaveLogic {
	return &SystemShopSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemShopSaveLogic) SystemShopSave(req *types.ShopSaveRequest) (resp *types.ShopSaveResponse, err error) {
	shop := types.GlobalShopConvertor.EntityToDomain(req.Shop)
	if err := transaction.UseTrans(l.ctx, l.svcCtx.DB, func(ctx context.Context, conn transaction.Conn) error {
		shop, err = l.svcCtx.ShopRepository.Insert(l.ctx, conn, shop)
		return err
	}, true); err != nil {
		return nil, err
	}
	resp = &types.ShopSaveResponse{}
	return
}
