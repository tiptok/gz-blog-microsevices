package skateboard

import (
	"github.com/tiptok/gz-blog-microsevices/pkg/result"
	"net/http"

	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/logic/skateboard"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SystemCoachCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CoachCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := skateboard.NewSystemCoachCreateLogic(r.Context(), svcCtx)
		resp, err := l.SystemCoachCreate(&req)
		result.HttpResult(r, w, resp, err)
	}
}
