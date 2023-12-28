package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminBindCourseCategoryLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	CategoryRepo repository.CategoryRepository
}

func NewAdminBindCourseCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminBindCourseCategoryLogic {
	categoryRepo := repository.NewCategoryRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminBindCourseCategoryLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		CategoryRepo: categoryRepo,
	}
}

func (l *AdminBindCourseCategoryLogic) AdminBindCourseCategory(req *types.AdminBindCourseCategoryReq) (resp *types.AdminBindCourseCategoryRes, err error) {
	resp, _ = l.CategoryRepo.AdminBindCourseCategory(l.ctx, req)

	return
}
