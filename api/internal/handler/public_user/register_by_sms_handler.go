package public_user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Sanagiig/fox-admin-core/api/internal/logic/public_user"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"
)

// swagger:route post /user/register_by_sms public_user RegisterBySms
//
// Register by SMS | 短信注册
//
// Register by SMS | 短信注册
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: RegisterBySmsReq
//
// Responses:
//  200: BaseMsgResp

func RegisterBySmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterBySmsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public_user.NewRegisterBySmsLogic(r.Context(), svcCtx)
		resp, err := l.RegisterBySms(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
