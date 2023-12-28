package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDeleteUserByIdLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	UserRepo repository.UserRepository
}

func NewAdminDeleteUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDeleteUserByIdLogic {
	userRepo := repository.NewUserRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminDeleteUserByIdLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		UserRepo: userRepo,
	}
}

func (l *AdminDeleteUserByIdLogic) AdminDeleteUserById(req *types.AdminDeleteUserByIdReq) (resp *types.AdminDeleteUserByIdRes, err error) {
	resp, _ = l.UserRepo.AdminDeleteUserById(l.ctx, req)
	return
}
