package student

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/internal/svc"
	"github.com/Sanagiig/fox-admin-core/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/types/core"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentLogic) UpdateStudent(in *core.StudentInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.Student.UpdateOneID(*in.Id).
			SetNotNilName(in.Name)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int(*in.Age)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: errormsg.UpdateSuccess }, nil
}
