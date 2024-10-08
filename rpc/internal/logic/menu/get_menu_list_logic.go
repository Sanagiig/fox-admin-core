package menu

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/ent/menu"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListLogic) GetMenuList(in *core.MenuListReq) (*core.MenuListResp, error) {
	var predicates []predicate.Menu
	if in.Path != nil {
		predicates = append(predicates, menu.PathContains(*in.Path))
	}
	if in.Name != nil {
		predicates = append(predicates, menu.NameContains(*in.Name))
	}
	if in.Redirect != nil {
		predicates = append(predicates, menu.RedirectContains(*in.Redirect))
	}
	result, err := l.svcCtx.DB.Menu.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MenuListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MenuInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Sort:	&v.Sort,
			ParentId:	&v.ParentID,
			MenuLevel:	&v.MenuLevel,
			MenuType:	&v.MenuType,
			Path:	&v.Path,
			Name:	&v.Name,
			Redirect:	&v.Redirect,
			Component:	&v.Component,
			Disabled:	&v.Disabled,
			ServiceName:	&v.ServiceName,
			Permission:	&v.Permission,
			Title:	&v.Title,
			Icon:	&v.Icon,
			HideMenu:	&v.HideMenu,
			HideBreadcrumb:	&v.HideBreadcrumb,
			IgnoreKeepAlive:	&v.IgnoreKeepAlive,
			HideTab:	&v.HideTab,
			FrameSrc:	&v.FrameSrc,
			CarryParam:	&v.CarryParam,
			HideChildrenInMenu:	&v.HideChildrenInMenu,
			Affix:	&v.Affix,
			DynamicLevel:	&v.DynamicLevel,
			RealPath:	&v.RealPath,
		})
	}

	return resp, nil
}
