package dictionary

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/ent/dictionary"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(in *core.DictionaryListReq) (*core.DictionaryListResp, error) {
	var predicates []predicate.Dictionary
	if in.Title != nil {
		predicates = append(predicates, dictionary.TitleContains(*in.Title))
	}
	if in.Name != nil {
		predicates = append(predicates, dictionary.NameContains(*in.Name))
	}
	if in.Desc != nil {
		predicates = append(predicates, dictionary.DescContains(*in.Desc))
	}
	result, err := l.svcCtx.DB.Dictionary.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.DictionaryListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.DictionaryInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Title:	&v.Title,
			Name:	&v.Name,
			Desc:	&v.Desc,
		})
	}

	return resp, nil
}
