package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdatePasswordLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	UserRepo repository.UserRepository
}

func NewUserUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdatePasswordLogic {
	userRepo := repository.NewUserRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &UserUpdatePasswordLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		UserRepo: userRepo,
	}
}

// UserUpdatePassword
// @Description 用户修改密码
// @Author Oberl-Fitzgerald 2023-12-27 14:10:44
// @Param  req *types.UserUpdatePasswordReq
// @Return resp
// @Return err
func (l *UserUpdatePasswordLogic) UserUpdatePassword(req *types.UserUpdatePasswordReq) (resp *types.UserUpdatePasswordRes, err error) {
	resp, _ = l.UserRepo.UserUpdatePassword(l.ctx, req)
	return
}
