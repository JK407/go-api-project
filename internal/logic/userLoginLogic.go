package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	UserRepo repository.UserRepository
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	userRepo := repository.NewUserRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &UserLoginLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		UserRepo: userRepo,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginRes, err error) {
	resp, _ = l.UserRepo.UserLogin(l.ctx, req)
	return
}
