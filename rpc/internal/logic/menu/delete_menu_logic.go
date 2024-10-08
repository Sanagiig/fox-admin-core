package menu

import (
	"context"

    "github.com/Sanagiig/fox-admin-core/rpc/ent/menu"
    "github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
    "github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/Sanagiig/fox-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Menu.Delete().Where(menu.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
