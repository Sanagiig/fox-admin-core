package publicapi

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicSystemConfigurationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicSystemConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicSystemConfigurationListLogic {
	return &GetPublicSystemConfigurationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPublicSystemConfigurationListLogic) GetPublicSystemConfigurationList() (resp *types.ConfigurationListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
