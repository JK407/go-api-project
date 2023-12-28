package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseListLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	CourseRepo repository.CourseRepository
}

func NewGetCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseListLogic {
	courseRepo := repository.NewCourseRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &GetCourseListLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		CourseRepo: courseRepo,
	}
}

func (l *GetCourseListLogic) GetCourseList(req *types.GetCourseListReq) (resp *types.GetCourseListRes, err error) {
	resp, _ = l.CourseRepo.GetCourseList(l.ctx, req)
	return
}
