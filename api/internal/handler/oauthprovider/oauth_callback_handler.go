package oauthprovider

import (
	"github.com/Sanagiig/fox-admin-core/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Sanagiig/fox-admin-core/api/internal/logic/oauthprovider"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
)

// swagger:route get /oauth/login/callback oauthprovider OauthCallback
//
// Oauth log in callback route | Oauth 登录返回调接口
//
// Oauth log in callback route | Oauth 登录返回调接口
//
// Responses:
//  200: CallbackResp

func OauthCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthProviderListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauthprovider.NewGetOauthProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetOauthProviderList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
