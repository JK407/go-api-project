package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDelCategoryLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	CategoryRepo repository.CategoryRepository
}

func NewAdminDelCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDelCategoryLogic {
	categoryRepo := repository.NewCategoryRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminDelCategoryLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		CategoryRepo: categoryRepo,
	}
}

func (l *AdminDelCategoryLogic) AdminDelCategory(req *types.AdminDelCategoryReq) (resp *types.AdminDelCategoryRes, err error) {
	resp, _ = l.CategoryRepo.AdminDelCategory(l.ctx, req)
	return
}
