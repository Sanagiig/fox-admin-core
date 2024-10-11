package base

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitJobDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitJobDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitJobDatabaseLogic {
	return &InitJobDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *InitJobDatabaseLogic) InitJobDatabase() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
