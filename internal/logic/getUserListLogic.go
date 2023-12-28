package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	UserRepo repository.UserRepository
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	userRepo := repository.NewUserRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &GetUserListLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		UserRepo: userRepo,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListRes, err error) {
	resp, _ = l.UserRepo.GetUserList(l.ctx, req)
	return
}
