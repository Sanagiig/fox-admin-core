package dictionary_detail

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/ent/dictionarydetail"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(in *core.DictionaryDetailListReq) (*core.DictionaryDetailListResp, error) {
	var predicates []predicate.DictionaryDetail
	if in.Title != nil {
		predicates = append(predicates, dictionarydetail.TitleContains(*in.Title))
	}
	if in.Key != nil {
		predicates = append(predicates, dictionarydetail.KeyContains(*in.Key))
	}
	if in.Value != nil {
		predicates = append(predicates, dictionarydetail.ValueContains(*in.Value))
	}
	result, err := l.svcCtx.DB.DictionaryDetail.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.DictionaryDetailListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.DictionaryDetailInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Sort:	&v.Sort,
			Title:	&v.Title,
			Key:	&v.Key,
			Value:	&v.Value,
			DictionaryId:	&v.DictionaryID,
		})
	}

	return resp, nil
}
