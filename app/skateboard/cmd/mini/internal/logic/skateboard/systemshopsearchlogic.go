package skateboard

import (
	"context"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/domain"
	"github.com/tiptok/gz-blog-microsevices/pkg/xerr"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemShopSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemShopSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemShopSearchLogic {
	return &SystemShopSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemShopSearchLogic) SystemShopSearch(req *types.ShopSearchRequest) (resp *types.ShopSearchResponse, err error) {
	conn := l.svcCtx.DefaultDBConn()
	queryOptions := domain.NewQueryOptions().WithOffsetLimit(req.Page, req.Size).WithKV("sortByRank", "desc")
	total, shops, err := l.svcCtx.ShopRepository.Find(l.ctx, conn, queryOptions)
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	list := make([]types.ShopItem, 0)
	for i := range shops {
		list = append(list, types.GlobalShopConvertor.DomainToEntity(shops[i]))
	}
	resp = &types.ShopSearchResponse{
		Total: total,
		List:  list,
	}
	return
}
