package student

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/internal/svc"
	"github.com/Sanagiig/fox-admin-core/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIdLogic {
	return &GetStudentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentByIdLogic) GetStudentById(in *core.IDReq) (*core.StudentInfo, error) {
	result, err := l.svcCtx.DB.Student.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.StudentInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Name:	&result.Name,
		Age:	pointy.GetPointer(int64(result.Age)),
	}, nil
}

