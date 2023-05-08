// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	skateboard "github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/handler/skateboard"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/cmd/mini/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: skateboard.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user-info",
				Handler: skateboard.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shop-list",
				Handler: skateboard.ShopListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/coach-list",
				Handler: skateboard.ShopCoachListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/coach-reserve",
				Handler: skateboard.CoachReserveHandler(serverCtx),
			},
		},
		rest.WithPrefix("/skateboard/mini"),
	)
}
