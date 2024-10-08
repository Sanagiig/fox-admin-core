package department

import (
	"context"

    "github.com/Sanagiig/fox-admin-core/rpc/ent/department"
    "github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
    "github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/Sanagiig/fox-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDepartmentLogic {
	return &DeleteDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDepartmentLogic) DeleteDepartment(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Department.Delete().Where(department.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
