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

type CreateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStudentLogic {
	return &CreateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStudentLogic) CreateStudent(in *core.StudentInfo) (*core.BaseIDResp, error) {
    query := l.svcCtx.DB.Student.Create().
			SetNotNilName(in.Name)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int(*in.Age)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: errormsg.CreateSuccess }, nil
}
