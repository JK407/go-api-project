package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateCategoryLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	CategoryRepo repository.CategoryRepository
}

func NewAdminUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateCategoryLogic {
	categoryRepo := repository.NewCategoryRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminUpdateCategoryLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		CategoryRepo: categoryRepo,
	}
}

func (l *AdminUpdateCategoryLogic) AdminUpdateCategory(req *types.AdminUpdateCategoryReq) (resp *types.AdminUpdateCategoryRes, err error) {
	resp, _ = l.CategoryRepo.AdminUpdateCategory(l.ctx, req)
	return
}
