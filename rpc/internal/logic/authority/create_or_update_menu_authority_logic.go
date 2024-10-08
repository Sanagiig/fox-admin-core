package authority

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(in *core.RoleMenuAuthorityReq) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseResp{}, nil
}
