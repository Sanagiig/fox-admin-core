package user

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *core.UserInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.User.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilUsername(in.Username).
			SetNotNilPassword(in.Password).
			SetNotNilNickname(in.Nickname).
			SetNotNilDescription(in.Description).
			SetNotNilHomePath(in.HomePath).
			SetNotNilMobile(in.Mobile).
			SetNotNilEmail(in.Email).
			SetNotNilAvatar(in.Avatar).
			SetNotNilDepartmentID(in.DepartmentId)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
