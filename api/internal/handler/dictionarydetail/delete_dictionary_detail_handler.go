package dictionarydetail

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Sanagiig/fox-admin-core/api/internal/logic/dictionarydetail"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"
)

// swagger:route post /dictionarydetail/delete dictionarydetail DeleteDictionaryDetail
//
// Delete dictionary detail information | 删除字典键值信息
//
// Delete dictionary detail information | 删除字典键值信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionarydetail.NewDeleteDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDictionaryDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
