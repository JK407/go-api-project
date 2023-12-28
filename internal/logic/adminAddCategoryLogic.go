package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminAddCategoryLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	CategoryRepo repository.CategoryRepository
}

func NewAdminAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminAddCategoryLogic {
	categoryRepo := repository.NewCategoryRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminAddCategoryLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		CategoryRepo: categoryRepo,
	}
}

func (l *AdminAddCategoryLogic) AdminAddCategory(req *types.AdminAddCategoryReq) (resp *types.AdminAddCategoryRes, err error) {
	resp, _ = l.CategoryRepo.AdminAddCategory(l.ctx, req)
	return
}
