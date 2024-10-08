package token

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/ent/token"
	"github.com/Sanagiig/fox-admin-core/rpc/ent/predicate"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *core.TokenListReq) (*core.TokenListResp, error) {
	var predicates []predicate.Token
	if in.Username != nil {
		predicates = append(predicates, token.UsernameContains(*in.Username))
	}
	if in.Token != nil {
		predicates = append(predicates, token.TokenContains(*in.Token))
	}
	if in.Source != nil {
		predicates = append(predicates, token.SourceContains(*in.Source))
	}
	result, err := l.svcCtx.DB.Token.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.TokenListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.TokenInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Uuid:	pointy.GetPointer(v.UUID.String()),
			Username:	&v.Username,
			Token:	&v.Token,
			Source:	&v.Source,
			ExpiredAt:	pointy.GetPointer(v.ExpiredAt.UnixMilli()),
		})
	}

	return resp, nil
}
