package authority

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(in *core.IDReq) (*core.RoleMenuAuthorityResp, error) {
	// todo: add your logic here and delete this line

	return &core.RoleMenuAuthorityResp{}, nil
}
