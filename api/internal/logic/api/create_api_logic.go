package api

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/internal/svc"
	"github.com/Sanagiig/fox-admin-core/internal/types"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiLogic) CreateApi(req *types.ApiInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.coreRpc.CreateApi(l.ctx,
		&__.ApiInfo{ 
        	CreatedAt: req.CreatedAt,
        	UpdatedAt: req.UpdatedAt,
        	Path: req.Path,
        	Description: req.Description,
        	ApiGroup: req.ApiGroup,
        	ServiceName: req.ServiceName,
        	Method: req.Method,
        	IsRequired: req.IsRequired,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
