package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_project/internal/constants"
	"go_project/internal/model/model"
	"go_project/internal/types"
	"go_project/internal/utils"
	"gorm.io/gorm"
)

type CourseRepository interface {
	AdminAddCourse(ctx context.Context, req *types.AdminAddCourseReq) (*types.AdminAddCourseRes, error)          // AdminAddCourse 管理员添加课程
	AdminUpdateCourse(ctx context.Context, req *types.AdminUpdateCourseReq) (*types.AdminUpdateCourseRes, error) // AdminUpdateCourse 管理员修改课程价格
	AdminDelCourse(ctx context.Context, req *types.AdminDelCourseReq) (*types.AdminDelCourseRes, error)          // AdminDelCourse 管理员删除课程
	GetCourseList(ctx context.Context, req *types.GetCourseListReq) (*types.GetCourseListRes, error)             // GetCourseList 获取课程列表
}

type CourseRepositoryImpl struct {
	CourseDB    *gorm.DB
	CourseRedis *redis.Client
}

// NewCourseRepository 初始化 CourseRepositoryImpl 结构体实例的函数
func NewCourseRepository(courseDB *gorm.DB, courseRedis *redis.Client) CourseRepository {
	return &CourseRepositoryImpl{
		CourseDB:    courseDB,
		CourseRedis: courseRedis,
	}
}

// AdminAddCourse
// @Description 管理员添加课程，课程名称能重复
// @Author Oberl-Fitzgerald 2023-12-28 10:22:04
// @Param  ctx context.Context
// @Param  req *types.AdminAddCourseReq
// @Return *types.AdminAddCourseRes
// @Return error
func (co *CourseRepositoryImpl) AdminAddCourse(ctx context.Context, req *types.AdminAddCourseReq) (*types.AdminAddCourseRes, error) {
	var user model.User
	//	根据手机号获取用户信息
	if findUserErr := co.CourseDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminAddCourseRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 判断用户密码是否正确
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminAddCourseRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 判断用户是否为管理员
	if user.RoleType != constants.AdminRole {
		return &types.AdminAddCourseRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 添加课程
	course := model.Course{
		CourseName:  req.CourseName,
		Description: req.Description,
		Price:       req.Price,
		UserID:      user.UserID,
	}
	if createCourseErr := co.CourseDB.WithContext(ctx).Create(&course).Error; createCourseErr != nil {
		return &types.AdminAddCourseRes{
			Code: constants.CourseAddErr.Code,
			Msg:  constants.CourseAddErr.Message,
		}, nil
	}
	return &types.AdminAddCourseRes{
		Code: constants.CourseAddSuc.Code,
		Msg:  constants.CourseAddSuc.Message,
		Data: &types.CourseData{
			CourseId:    course.CourseID,
			UserId:      user.UserID,
			CourseName:  course.CourseName,
			Description: course.Description,
			Price:       course.Price,
		},
	}, nil

}

// AdminUpdateCourse
// @Description 管理员修改课程价格
// @Author Oberl-Fitzgerald 2023-12-28 10:59:53
// @Param  ctx context.Context
// @Param  req *types.AdminUpdateCourseReq
// @Return *types.AdminUpdateCourseRes
// @Return error
func (co *CourseRepositoryImpl) AdminUpdateCourse(ctx context.Context, req *types.AdminUpdateCourseReq) (*types.AdminUpdateCourseRes, error) {
	var user model.User
	//	根据手机号获取用户信息
	if findUserErr := co.CourseDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminUpdateCourseRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 判断用户密码是否正确
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminUpdateCourseRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 判断用户是否为管理员
	if user.RoleType != constants.AdminRole {
		return &types.AdminUpdateCourseRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 根据课程id查询课程
	var courseById model.Course
	if findCourseErr := co.CourseDB.WithContext(ctx).Where("course_id = ? and status = ? ", req.CourseId, constants.StatusNormal).First(&courseById).Error; findCourseErr != nil {
		return &types.AdminUpdateCourseRes{
			Code: constants.CourseNotExist.Code,
			Msg:  constants.CourseNotExist.Message,
		}, nil
	}
	// 更新课程价格和用户id
	if updateCourseErr := co.CourseDB.WithContext(ctx).Model(&courseById).Updates(model.Course{Price: req.Price, UserID: user.UserID}).Error; updateCourseErr != nil {
		return &types.AdminUpdateCourseRes{
			Code: constants.CourseUpdateErr.Code,
			Msg:  constants.CourseUpdateErr.Message,
		}, nil
	}
	return &types.AdminUpdateCourseRes{
		Code: constants.CourseUpdateSuc.Code,
		Msg:  constants.CourseUpdateSuc.Message,
		Data: &types.CourseData{
			CourseId:    courseById.CourseID,
			UserId:      user.UserID,
			CourseName:  courseById.CourseName,
			Description: courseById.Description,
			Price:       courseById.Price,
		},
	}, nil

}

// AdminDelCourse
// @Description 管理员删除课程
// @Author Oberl-Fitzgerald 2023-12-28 11:53:45
// @Param  ctx context.Context
// @Param  req *types.AdminDelCourseReq
// @Return *types.AdminDelCourseRes
// @Return error
func (co *CourseRepositoryImpl) AdminDelCourse(ctx context.Context, req *types.AdminDelCourseReq) (*types.AdminDelCourseRes, error) {
	var user model.User
	//	根据手机号获取用户信息
	if findUserErr := co.CourseDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminDelCourseRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 判断用户密码是否正确
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminDelCourseRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 判断用户是否为管理员
	if user.RoleType != constants.AdminRole {
		return &types.AdminDelCourseRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 根据课程id查询课程
	var courseById model.Course
	if findCourseErr := co.CourseDB.WithContext(ctx).Where("course_id = ? and status = ? ", req.CourseId, constants.StatusNormal).First(&courseById).Error; findCourseErr != nil {
		return &types.AdminDelCourseRes{
			Code: constants.CourseNotExist.Code,
			Msg:  constants.CourseNotExist.Message,
		}, nil
	}
	// 开启事务
	tx := co.CourseDB.WithContext(ctx).Begin()
	// 删除课程和用户id
	if updateCourseErr := tx.Model(&courseById).Updates(model.Course{UserID: user.UserID, Status: constants.StatusBan}).Error; updateCourseErr != nil {
		// 回滚事务
		tx.Rollback()
		return &types.AdminDelCourseRes{
			Code: constants.CourseDeleteErr.Code,
			Msg:  constants.CourseDeleteErr.Message,
		}, nil
	}
	// 删除关系表中的对应课程id的绑定数据，假删除
	var courseCategoryShip model.CourseCategoryShip
	if updateCourseCategoryErr := tx.Model(&courseCategoryShip).Where("course_id = ? and status = ? ", req.CourseId, constants.StatusNormal).Updates(model.CourseCategory{Status: constants.StatusBan}).Error; updateCourseCategoryErr != nil {
		// 回滚事务
		tx.Rollback()
		return &types.AdminDelCourseRes{
			Code: constants.CourseDeleteErr.Code,
			Msg:  constants.CourseDeleteErr.Message,
		}, nil
	}
	// 提交事务
	tx.Commit()
	return &types.AdminDelCourseRes{
		Code: constants.CourseDeleteSuc.Code,
		Msg:  constants.CourseDeleteSuc.Message,
	}, nil
}

// GetCourseList
// @Description 获取课程列表,分页查询，支持按照课程名查询
// @Author Oberl-Fitzgerald 2023-12-28 14:10:48
// @Param  ctx context.Context
// @Param  req *types.GetCourseListReq
// @Return *types.GetCourseListRes
// @Return error
func (co *CourseRepositoryImpl) GetCourseList(ctx context.Context, req *types.GetCourseListReq) (*types.GetCourseListRes, error) {
	var CourseList []*model.Course
	var total int64
	// 计算偏移量
	offset := (req.Page - 1) * req.Size
	// 查询课程列表
	if findCourseErr := co.CourseDB.WithContext(ctx).Where("status = ? and course_name like ?", constants.StatusNormal, "%"+req.Keyword+"%").Offset(offset).Limit(req.Size).Find(&CourseList).Error; findCourseErr != nil {
		return &types.GetCourseListRes{
			Code: constants.CourseListErr.Code,
			Msg:  constants.CourseListErr.Message,
		}, nil
	}
	// 查询课程总数
	if findCourseCountErr := co.CourseDB.WithContext(ctx).Model(&model.Course{}).Where("status = ? and course_name like ? ", constants.StatusNormal, "%"+req.Keyword+"%").Count(&total).Error; findCourseCountErr != nil {
		return &types.GetCourseListRes{
			Code: constants.CourseCountErr.Code,
			Msg:  constants.CourseCountErr.Message,
		}, nil
	}
	// 将课程列表转换为返回数据
	var courseDataList []*types.CourseData
	for _, course := range CourseList {
		courseData := &types.CourseData{
			CourseId:    course.CourseID,
			UserId:      course.UserID,
			CourseName:  course.CourseName,
			Description: course.Description,
			Price:       course.Price,
		}
		courseDataList = append(courseDataList, courseData)
	}
	return &types.GetCourseListRes{
		Code:  constants.CourseListSuc.Code,
		Msg:   constants.CourseListSuc.Message,
		Data:  courseDataList,
		Total: int(total),
	}, nil
}
