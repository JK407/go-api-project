package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminAddCourseLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	CourseRepo repository.CourseRepository // 将字段类型改为接口类型
}

func NewAdminAddCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminAddCourseLogic {
	// 初始化 CourseRepositoryImpl 结构体实例，并转换为接口类型
	courseRepo := repository.NewCourseRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &AdminAddCourseLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		CourseRepo: courseRepo, // 使用初始化后的接口实例
	}
}

func (l *AdminAddCourseLogic) AdminAddCourse(req *types.AdminAddCourseReq) (resp *types.AdminAddCourseRes, err error) {
	resp, _ = l.CourseRepo.AdminAddCourse(l.ctx, req)
	return
}
