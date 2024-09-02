package student

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/ent/student"
	"github.com/Sanagiig/fox-admin-core/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/internal/svc"
	"github.com/Sanagiig/fox-admin-core/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentListLogic) GetStudentList(in *core.StudentListReq) (*core.StudentListResp, error) {
	var predicates []predicate.Student
	if in.Name != nil {
		predicates = append(predicates, student.NameContains(*in.Name))
	}
	result, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.StudentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.StudentInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:	&v.Name,
			Age:	pointy.GetPointer(int64(v.Age)),
		})
	}

	return resp, nil
}
