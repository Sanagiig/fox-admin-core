package student

import (
	"context"

    "github.com/Sanagiig/fox-admin-core/ent/student"
    "github.com/Sanagiig/fox-admin-core/internal/svc"
    "github.com/Sanagiig/fox-admin-core/internal/utils/dberrorhandler"
    "github.com/Sanagiig/fox-admin-core/types/core"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentLogic {
	return &DeleteStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStudentLogic) DeleteStudent(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Student.Delete().Where(student.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: errormsg.DeleteSuccess }, nil
}
