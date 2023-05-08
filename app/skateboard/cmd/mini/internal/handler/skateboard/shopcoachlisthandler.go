package skateboard

import (
	"net/http"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/logic/skateboard"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShopCoachListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CoachListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := skateboard.NewShopCoachListLogic(r.Context(), svcCtx)
		resp, err := l.ShopCoachList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
