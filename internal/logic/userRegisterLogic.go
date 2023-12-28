package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	UserRepo repository.UserRepository
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	userRepo := repository.NewUserRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &UserRegisterLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		UserRepo: userRepo,
	}
}

// UserRegister
// @Description 用户注册
// @Author Oberl-Fitzgerald 2023-12-27 09:33:03
// @Param  req *types.UserRegisterReq
// @Return resp
// @Return err
func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterRes, err error) {
	resp, _ = l.UserRepo.UserRegister(l.ctx, req)
	return
}
