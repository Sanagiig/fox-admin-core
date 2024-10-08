package api

import (
    "context"

    "github.com/Sanagiig/fox-admin-core/internal/svc"
    "github.com/Sanagiig/fox-admin-core/internal/types"
    "github.com/Sanagiig/fox-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiByIdLogic) GetApiById(req *types.IDReq) (resp *types.ApiInfoResp, err error) {
	data, err := l.svcCtx.coreRpc.GetApiById(l.ctx, &__.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.ApiInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ApiInfo{
			Id:  data.Id,
        	CreatedAt: data.CreatedAt,
        	UpdatedAt: data.UpdatedAt,
        	Path: data.Path,
        	Description: data.Description,
        	ApiGroup: data.ApiGroup,
        	ServiceName: data.ServiceName,
        	Method: data.Method,
        	IsRequired: data.IsRequired,
		},
	}, nil
}

