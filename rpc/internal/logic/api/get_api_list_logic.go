package api

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/ent/api"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListLogic) GetApiList(in *core.ApiListReq) (*core.ApiListResp, error) {
	var predicates []predicate.Api
	if in.Path != nil {
		predicates = append(predicates, api.PathContains(*in.Path))
	}
	if in.Description != nil {
		predicates = append(predicates, api.DescriptionContains(*in.Description))
	}
	if in.ApiGroup != nil {
		predicates = append(predicates, api.APIGroupContains(*in.ApiGroup))
	}
	result, err := l.svcCtx.DB.Api.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.ApiListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.ApiInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Path:	&v.Path,
			Description:	&v.Description,
			ApiGroup:	&v.APIGroup,
			ServiceName:	&v.ServiceName,
			Method:	&v.Method,
			IsRequired:	&v.IsRequired,
		})
	}

	return resp, nil
}
