package menu

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuByIdLogic {
	return &GetMenuByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuByIdLogic) GetMenuById(in *core.IDReq) (*core.MenuInfo, error) {
	result, err := l.svcCtx.DB.Menu.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.MenuInfo{
		Id:          &result.ID,
		CreatedAt:   pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:   pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Sort:        &result.Sort,
		ParentId:    &result.ParentID,
		MenuType:    &result.MenuType,
		Path:        &result.Path,
		Name:        &result.Name,
		Redirect:    &result.Redirect,
		Component:   &result.Component,
		Disabled:    &result.Disabled,
		ServiceName: &result.ServiceName,
		Permission:  &result.Permission,
		Level:       &result.MenuLevel,
		Meta: &core.Meta{
			Title:              &result.Title,
			Icon:               &result.Icon,
			HideMenu:           &result.HideMenu,
			HideBreadcrumb:     &result.HideBreadcrumb,
			IgnoreKeepAlive:    &result.IgnoreKeepAlive,
			HideTab:            &result.HideTab,
			FrameSrc:           &result.FrameSrc,
			CarryParam:         &result.CarryParam,
			HideChildrenInMenu: &result.HideChildrenInMenu,
			Affix:              &result.Affix,
			DynamicLevel:       &result.DynamicLevel,
			RealPath:           &result.RealPath,
		},
	}, nil
}
