// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"github.com/Sanagiig/fox-admin-core/rpc/internal/logic/base"
	"github.com/Sanagiig/fox-admin-core/rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/rpc/types/core"
)

type CoreServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedCoreServer
}

func NewCoreServer(svcCtx *svc.ServiceContext) *CoreServer {
	return &CoreServer{
		svcCtx: svcCtx,
	}
}

func (s *CoreServer) InitDatabase(ctx context.Context, in *core.Empty) (*core.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}
