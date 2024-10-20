package dictionarydetail

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Sanagiig/fox-admin-core/api/internal/logic/dictionarydetail"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"
)

// swagger:route post /dictionarydetail/list dictionarydetail GetDictionaryDetailList
//
// Get dictionary detail list | 获取字典键值列表
//
// Get dictionary detail list | 获取字典键值列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryDetailListReq
//
// Responses:
//  200: DictionaryDetailListResp

func GetDictionaryDetailListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryDetailListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionarydetail.NewGetDictionaryDetailListLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryDetailList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
