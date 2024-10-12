package menu

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
