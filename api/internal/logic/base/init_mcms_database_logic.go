package base

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitMcmsDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitMcmsDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitMcmsDatabaseLogic {
	return &InitMcmsDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *InitMcmsDatabaseLogic) InitMcmsDatabase() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
