package menu

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *core.MenuInfo) (*core.BaseResp, error) {
	err:= l.svcCtx.DB.Menu.UpdateOneID(*in.Id).
			SetNotNilSort(in.Sort).
			SetNotNilParentID(in.ParentId).
			SetNotNilMenuLevel(in.MenuLevel).
			SetNotNilMenuType(in.MenuType).
			SetNotNilPath(in.Path).
			SetNotNilName(in.Name).
			SetNotNilRedirect(in.Redirect).
			SetNotNilComponent(in.Component).
			SetNotNilDisabled(in.Disabled).
			SetNotNilServiceName(in.ServiceName).
			SetNotNilPermission(in.Permission).
			SetNotNilTitle(in.Title).
			SetNotNilIcon(in.Icon).
			SetNotNilHideMenu(in.HideMenu).
			SetNotNilHideBreadcrumb(in.HideBreadcrumb).
			SetNotNilIgnoreKeepAlive(in.IgnoreKeepAlive).
			SetNotNilHideTab(in.HideTab).
			SetNotNilFrameSrc(in.FrameSrc).
			SetNotNilCarryParam(in.CarryParam).
			SetNotNilHideChildrenInMenu(in.HideChildrenInMenu).
			SetNotNilAffix(in.Affix).
			SetNotNilDynamicLevel(in.DynamicLevel).
			SetNotNilRealPath(in.RealPath).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
