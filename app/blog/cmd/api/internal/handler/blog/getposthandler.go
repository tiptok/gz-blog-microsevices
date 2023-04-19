package blog

import (
	"net/http"

	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/logic/blog"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/svc"
	"github.com/tiptok/gz-blog-microsevices/app/blog/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := blog.NewGetPostLogic(r.Context(), svcCtx)
		resp, err := l.GetPost(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
