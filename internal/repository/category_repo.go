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

type CategoryRepository interface {
	AdminAddCategory(ctx context.Context, req *types.AdminAddCategoryReq) (*types.AdminAddCategoryRes, error)                      // AdminAddCategory 管理员添加课程
	AdminUpdateCategory(ctx context.Context, req *types.AdminUpdateCategoryReq) (*types.AdminUpdateCategoryRes, error)             // AdminUpdateCategory 管理员修改课程名称
	AdminBindCourseCategory(ctx context.Context, req *types.AdminBindCourseCategoryReq) (*types.AdminBindCourseCategoryRes, error) // AdminBindCourseCategory 管理员绑定课程分类
	AdminDelCategory(ctx context.Context, req *types.AdminDelCategoryReq) (*types.AdminDelCategoryRes, error)                      // AdminDelCategory 管理员删除课程分类
	GetCategoryList(ctx context.Context, req *types.GetCategoryListReq) (*types.GetCategoryListRes, error)                         // GetCategoryList 获取课程分类列表
}

type CategoryRepositoryImpl struct {
	CategoryDB    *gorm.DB
	CategoryRedis *redis.Client
}

// GetCategoryList
// @Description 获取课程分类列表
// @Author Oberl-Fitzgerald 2023-12-28 16:17:51
// @Param  ctx context.Context
// @Param  req *types.GetCategoryListReq
// @Return *types.GetCategoryListRes
// @Return error
func (ca CategoryRepositoryImpl) GetCategoryList(ctx context.Context, req *types.GetCategoryListReq) (*types.GetCategoryListRes, error) {
	var categoryList []*model.CourseCategory
	var total int64
	// 计算偏移量
	offset := (req.Page - 1) * req.Size
	// 查询分类列表
	if findCategoryListErr := ca.CategoryDB.WithContext(ctx).Where("status = ? and category_name like ? ", constants.StatusNormal, "%"+req.Keyword+"%").Offset(offset).Limit(req.Size).Find(&categoryList).Error; findCategoryListErr != nil {
		return &types.GetCategoryListRes{
			Code: constants.CategoryListErr.Code,
			Msg:  constants.CategoryListErr.Message,
		}, nil
	}
	// 查询分类总数
	if findCategoryCountErr := ca.CategoryDB.WithContext(ctx).Model(&model.CourseCategory{}).Where("status = ? and category_name like ? ", constants.StatusNormal, "%"+req.Keyword+"%").Count(&total).Error; findCategoryCountErr != nil {
		return &types.GetCategoryListRes{
			Code: constants.CategoryCountErr.Code,
			Msg:  constants.CategoryCountErr.Message,
		}, nil
	}
	// 封装返回数据
	var categoryDataList []*types.CategoryData
	for _, category := range categoryList {
		categoryData := &types.CategoryData{
			CategoryId:   category.CategoryID,
			CategoryName: category.CategoryName,
		}
		categoryDataList = append(categoryDataList, categoryData)
	}
	return &types.GetCategoryListRes{
		Code:  constants.CategoryListSuc.Code,
		Msg:   constants.CategoryListSuc.Message,
		Data:  categoryDataList,
		Total: int(total),
	}, nil
}

// AdminDelCategory
// @Description 管理员删除课程分类
// @Author Oberl-Fitzgerald 2023-12-28 16:03:55
// @Param  ctx context.Context
// @Param  req *types.AdminDelCategoryReq
// @Return *types.AdminDelCategoryRes
// @Return error
func (ca CategoryRepositoryImpl) AdminDelCategory(ctx context.Context, req *types.AdminDelCategoryReq) (*types.AdminDelCategoryRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ca.CategoryDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminDelCategoryRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminDelCategoryRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 校验身份
	if user.RoleType != constants.AdminRole {
		return &types.AdminDelCategoryRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 根据分类id查询分类是否存在
	var category model.CourseCategory
	if findCategoryErr := ca.CategoryDB.WithContext(ctx).Where("category_id = ? and status = ? ", req.CategoryId, constants.StatusNormal).First(&category).Error; findCategoryErr != nil {
		// 分类不存在
		if findCategoryErr == gorm.ErrRecordNotFound {
			return &types.AdminDelCategoryRes{
				Code: constants.CategoryNotExist.Code,
				Msg:  constants.CategoryNotExist.Message,
			}, nil
		} else {
			// 查询分类失败
			return &types.AdminDelCategoryRes{
				Code: constants.ServerError.Code,
				Msg:  constants.ServerError.Message,
			}, nil
		}
	}
	// 开启事务
	tx := ca.CategoryDB.WithContext(ctx).Begin()
	// 删除分类
	if deleteCategoryErr := tx.Model(&category).Update("status", constants.StatusBan).Error; deleteCategoryErr != nil {
		tx.Rollback()
		return &types.AdminDelCategoryRes{
			Code: constants.CategoryDeleteErr.Code,
			Msg:  constants.CategoryDeleteErr.Message,
		}, nil
	}
	// 删除课程和分类的绑定关系
	if deleteCourseCategoryErr := tx.Model(&model.CourseCategoryShip{}).Where("category_id = ? and status = ? ", req.CategoryId, constants.StatusNormal).Update("status", constants.StatusBan).Error; deleteCourseCategoryErr != nil {
		tx.Rollback()
		return &types.AdminDelCategoryRes{
			Code: constants.CategoryDeleteErr.Code,
			Msg:  constants.CategoryDeleteErr.Message,
		}, nil
	}
	tx.Commit()
	return &types.AdminDelCategoryRes{
		Code: constants.CategoryDeleteSuc.Code,
		Msg:  constants.CategoryDeleteSuc.Message,
	}, nil
}

// AdminBindCourseCategory
// @Description 管理员绑定课程分类
// @Author Oberl-Fitzgerald 2023-12-28 16:03:51
// @Param  ctx context.Context
// @Param  req *types.AdminBindCourseCategoryReq
// @Return *types.AdminBindCourseCategoryRes
// @Return error
func (ca CategoryRepositoryImpl) AdminBindCourseCategory(ctx context.Context, req *types.AdminBindCourseCategoryReq) (*types.AdminBindCourseCategoryRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ca.CategoryDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminBindCourseCategoryRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminBindCourseCategoryRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 校验身份
	if user.RoleType != constants.AdminRole {
		return &types.AdminBindCourseCategoryRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 根据课程id查询课程是否存在
	var course model.Course
	if findCourseErr := ca.CategoryDB.WithContext(ctx).Where("course_id = ? and status = ? ", req.CourseId, constants.StatusNormal).First(&course).Error; findCourseErr != nil {
		// 课程不存在
		if findCourseErr == gorm.ErrRecordNotFound {
			return &types.AdminBindCourseCategoryRes{
				Code: constants.CourseNotExist.Code,
				Msg:  constants.CourseNotExist.Message,
			}, nil
		} else {
			// 查询课程失败
			return &types.AdminBindCourseCategoryRes{
				Code: constants.ServerError.Code,
				Msg:  constants.ServerError.Message,
			}, nil
		}
	}
	// 根据分类id查询分类是否存在
	var category model.CourseCategory
	if findCategoryErr := ca.CategoryDB.WithContext(ctx).Where("category_id = ? and status = ? ", req.CategoryId, constants.StatusNormal).First(&category).Error; findCategoryErr != nil {
		// 分类不存在
		if findCategoryErr == gorm.ErrRecordNotFound {
			return &types.AdminBindCourseCategoryRes{
				Code: constants.CategoryNotExist.Code,
				Msg:  constants.CategoryNotExist.Message,
			}, nil
		} else {
			// 查询分类失败
			return &types.AdminBindCourseCategoryRes{
				Code: constants.ServerError.Code,
				Msg:  constants.ServerError.Message,
			}, nil
		}
	}
	// 根据课程id和分类id查询是否已经绑定
	var courseCategoryShip model.CourseCategoryShip
	if findCourseCategoryErr := ca.CategoryDB.WithContext(ctx).Where("course_id = ? and category_id = ? and status = ? ", req.CourseId, req.CategoryId, constants.StatusNormal).First(&courseCategoryShip).Error; findCourseCategoryErr != nil {
		// 未绑定
		if findCourseCategoryErr == gorm.ErrRecordNotFound {
			// 绑定课程和分类
			courseCategory := model.CourseCategoryShip{
				CourseID:   int64(req.CourseId),
				CategoryID: int64(req.CategoryId),
			}
			if createCourseCategoryErr := ca.CategoryDB.WithContext(ctx).Create(&courseCategory).Error; createCourseCategoryErr != nil {
				return &types.AdminBindCourseCategoryRes{
					Code: constants.CourseCategoryBindErr.Code,
					Msg:  constants.CourseCategoryBindErr.Message,
				}, nil
			}
			return &types.AdminBindCourseCategoryRes{
				Code: constants.CourseCategoryBindSuc.Code,
				Msg:  constants.CourseCategoryBindSuc.Message,
				Data: &types.CourseCategoryData{
					Id:         courseCategory.ID,
					CourseId:   courseCategory.CourseID,
					CategoryId: courseCategory.CategoryID,
				},
			}, nil
		} else {
			// 查询绑定关系失败
			return &types.AdminBindCourseCategoryRes{
				Code: constants.ServerError.Code,
				Msg:  constants.ServerError.Message,
			}, nil
		}
	} else {
		// 已经绑定
		return &types.AdminBindCourseCategoryRes{
			Code: constants.CourseCategoryBindExist.Code,
			Msg:  constants.CourseCategoryBindExist.Message,
		}, nil
	}
}

// AdminUpdateCategory
// @Description 管理员修改课程分类名称
// @Author Oberl-Fitzgerald 2023-12-28 15:08:42
// @Param  ctx context.Context
// @Param  req *types.AdminUpdateCategoryReq
// @Return *types.AdminUpdateCategoryRes
// @Return error
func (ca CategoryRepositoryImpl) AdminUpdateCategory(ctx context.Context, req *types.AdminUpdateCategoryReq) (*types.AdminUpdateCategoryRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ca.CategoryDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminUpdateCategoryRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminUpdateCategoryRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 校验身份
	if user.RoleType != constants.AdminRole {
		return &types.AdminUpdateCategoryRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 根据校验分类是否存在
	var category model.CourseCategory
	if findCategoryByIdErr := ca.CategoryDB.WithContext(ctx).Where("category_id = ? and status = ? ", req.CategoryId, constants.StatusNormal).First(&category).Error; findCategoryByIdErr != nil {
		// 分类不存在
		if findCategoryByIdErr == gorm.ErrRecordNotFound {
			return &types.AdminUpdateCategoryRes{
				Code: constants.CategoryNotExist.Code,
				Msg:  constants.CategoryNotExist.Message,
			}, nil
		} else {
			// 查询分类失败
			return &types.AdminUpdateCategoryRes{
				Code: constants.ServerError.Code,
				Msg:  constants.ServerError.Message,
			}, nil
		}
	} else {
		// 分类存在
		// 修改分类名称
		// 通过名称查询分类是否存在
		var categoryByName model.CourseCategory
		if findCategoryByNameErr := ca.CategoryDB.WithContext(ctx).Where("category_name = ? and status = ? ", req.CategoryName, constants.StatusNormal).First(&categoryByName).Error; findCategoryByNameErr != nil {
			// 分类不存在,修改分类名称
			if findCategoryByNameErr == gorm.ErrRecordNotFound {
				if updateCategoryErr := ca.CategoryDB.WithContext(ctx).Model(&category).Update("category_name", req.CategoryName).Error; updateCategoryErr != nil {
					return &types.AdminUpdateCategoryRes{
						Code: constants.CategoryCreateErr.Code,
						Msg:  constants.CategoryCreateErr.Message,
					}, nil
				}
				return &types.AdminUpdateCategoryRes{
					Code: constants.CategoryCreateSuc.Code,
					Msg:  constants.CategoryCreateSuc.Message,
					Data: &types.CategoryData{
						CategoryId:   category.CategoryID,
						CategoryName: category.CategoryName,
					},
				}, nil
			} else {
				// 查询分类失败
				return &types.AdminUpdateCategoryRes{
					Code: constants.ServerError.Code,
					Msg:  constants.ServerError.Message,
				}, nil
			}
		} else {
			// 分类名称已存在
			return &types.AdminUpdateCategoryRes{
				Code: constants.CategoryNameExist.Code,
				Msg:  constants.CategoryNameExist.Message,
			}, nil
		}
	}
}

// AdminAddCategory
// @Description 管理员添加课程分类，分类名称不能重复
// @Author Oberl-Fitzgerald 2023-12-28 14:34:35
// @Param  ctx context.Context
// @Param  req *types.AdminAddCategoryReq
// @Return *types.AdminAddCategoryRes
// @Return error
func (ca CategoryRepositoryImpl) AdminAddCategory(ctx context.Context, req *types.AdminAddCategoryReq) (*types.AdminAddCategoryRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ca.CategoryDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminAddCategoryRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminAddCategoryRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 校验身份
	if user.RoleType != constants.AdminRole {
		return &types.AdminAddCategoryRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 通过名称校验分类是否存在
	var categoryByName model.CourseCategory
	if findCategoryErr := ca.CategoryDB.WithContext(ctx).Where("category_name = ? and status = ? ", req.CategoryName, constants.StatusNormal).First(&categoryByName).Error; findCategoryErr != nil {
		// 分类不存在,添加分类
		if findCategoryErr == gorm.ErrRecordNotFound {
			// 添加分类
			category := model.CourseCategory{
				CategoryName: req.CategoryName,
			}
			if createCategoryErr := ca.CategoryDB.WithContext(ctx).Create(&category).Error; createCategoryErr != nil {
				return &types.AdminAddCategoryRes{
					Code: constants.CategoryCreateErr.Code,
					Msg:  constants.CategoryCreateErr.Message,
				}, nil
			}
			return &types.AdminAddCategoryRes{
				Code: constants.CategoryCreateSuc.Code,
				Msg:  constants.CategoryCreateSuc.Message,
				Data: &types.CategoryData{
					CategoryId:   category.CategoryID,
					CategoryName: category.CategoryName,
				},
			}, nil
		} else {
			// 查询分类失败
			return &types.AdminAddCategoryRes{
				Code: constants.ServerError.Code,
				Msg:  constants.ServerError.Message,
			}, nil
		}
	} else {
		// 分类名称已存在
		return &types.AdminAddCategoryRes{
			Code: constants.CategoryNameExist.Code,
			Msg:  constants.CategoryNameExist.Message,
		}, nil
	}
}

// NewCategoryRepository 初始化 CategoryRepositoryImpl 结构体实例的函数
func NewCategoryRepository(categoryDB *gorm.DB, categoryRedis *redis.Client) CategoryRepository {
	return &CategoryRepositoryImpl{
		CategoryDB:    categoryDB,
		CategoryRedis: categoryRedis,
	}
}
