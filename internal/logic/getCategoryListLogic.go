package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	CategoryRepo repository.CategoryRepository
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	categoryRepo := repository.NewCategoryRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &GetCategoryListLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		CategoryRepo: categoryRepo,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.GetCategoryListReq) (resp *types.GetCategoryListRes, err error) {
	resp, _ = l.CategoryRepo.GetCategoryList(l.ctx, req)
	return
}
