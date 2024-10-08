package user

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/ent/user"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *core.UserListReq) (*core.UserListResp, error) {
	var predicates []predicate.User
	if in.Username != nil {
		predicates = append(predicates, user.UsernameContains(*in.Username))
	}
	if in.Password != nil {
		predicates = append(predicates, user.PasswordContains(*in.Password))
	}
	if in.Nickname != nil {
		predicates = append(predicates, user.NicknameContains(*in.Nickname))
	}
	result, err := l.svcCtx.DB.User.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.UserListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.UserInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Username:	&v.Username,
			Password:	&v.Password,
			Nickname:	&v.Nickname,
			Description:	&v.Description,
			HomePath:	&v.HomePath,
			Mobile:	&v.Mobile,
			Email:	&v.Email,
			Avatar:	&v.Avatar,
			DepartmentId:	&v.DepartmentID,
		})
	}

	return resp, nil
}
