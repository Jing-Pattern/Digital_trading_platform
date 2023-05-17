package handler

import (
	"Digital_trading_platform/common"
	"net/http"

	"Digital_trading_platform/service/user/api/internal/logic"
	"Digital_trading_platform/service/user/api/internal/svc"
	"Digital_trading_platform/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result := common.NewResult().Deal(resp, err)
		httpx.OkJson(w, result)
	}
}
