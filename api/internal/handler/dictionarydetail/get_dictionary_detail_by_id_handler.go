package dictionarydetail

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/Sanagiig/fox-admin-core/api/internal/logic/dictionarydetail"
	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"
)

// swagger:route post /dictionarydetail dictionarydetail GetDictionaryDetailById
//
// Get dictionary detail by ID | 通过ID获取字典键值
//
// Get dictionary detail by ID | 通过ID获取字典键值
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: DictionaryDetailInfoResp

func GetDictionaryDetailByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionarydetail.NewGetDictionaryDetailByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryDetailById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
