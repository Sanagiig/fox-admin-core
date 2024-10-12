package captcha

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/api/internal/svc"
	"github.com/Sanagiig/fox-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsCaptchaLogic {
	return &GetSmsCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSmsCaptchaLogic) GetSmsCaptcha(req *types.SmsCaptchaReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
