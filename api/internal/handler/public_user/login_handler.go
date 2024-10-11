package public_user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Sanagiig/fox-admin-core/api/internal/logic/public_user"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"
)

// swagger:route post /user/login public_user Login
//
// Log in | 登录
//
// Log in | 登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: LoginReq
//
// Responses:
//  200: LoginResp

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public_user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
