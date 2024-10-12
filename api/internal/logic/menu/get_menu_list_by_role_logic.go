package menu

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole() (resp *types.MenuListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
