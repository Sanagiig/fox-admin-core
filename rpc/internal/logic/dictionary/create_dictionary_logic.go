package dictionary

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDictionaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDictionaryLogic {
	return &CreateDictionaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDictionaryLogic) CreateDictionary(in *core.DictionaryInfo) (*core.BaseIDResp, error) {
    query := l.svcCtx.DB.Dictionary.Create().
			SetNotNilTitle(in.Title).
			SetNotNilName(in.Name).
			SetNotNilDesc(in.Desc)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
