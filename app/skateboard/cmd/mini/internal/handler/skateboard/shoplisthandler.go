package skateboard

import (
	"net/http"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/logic/skateboard"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShopListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShopListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := skateboard.NewShopListLogic(r.Context(), svcCtx)
		resp, err := l.ShopList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
