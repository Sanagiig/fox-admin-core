package api

import (
	"context"

    "github.com/Sanagiig/fox-admin-core/internal/svc"
	"github.com/Sanagiig/fox-admin-core/internal/types"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	data, err := l.svcCtx.coreRpc.GetApiList(l.ctx,
		&__.ApiListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			PageSize: req.PageSize,
			Path: req.Path,
			Description: req.Description,
			ApiGroup: req.ApiGroup,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ApiListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.ApiInfo{
			Id:  v.Id,
        	CreatedAt: v.CreatedAt,
        	UpdatedAt: v.UpdatedAt,
        	Path: v.Path,
        	Description: v.Description,
        	ApiGroup: v.ApiGroup,
        	ServiceName: v.ServiceName,
        	Method: v.Method,
        	IsRequired: v.IsRequired,
			})
	}
	return resp, nil
}
