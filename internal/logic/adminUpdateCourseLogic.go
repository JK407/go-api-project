package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateCourseLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	CourseRepo repository.CourseRepository
}

func NewAdminUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateCourseLogic {
	courseRepo := repository.NewCourseRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminUpdateCourseLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		CourseRepo: courseRepo,
	}
}

func (l *AdminUpdateCourseLogic) AdminUpdateCourse(req *types.AdminUpdateCourseReq) (resp *types.AdminUpdateCourseRes, err error) {
	resp, _ = l.CourseRepo.AdminUpdateCourse(l.ctx, req)
	return
}
