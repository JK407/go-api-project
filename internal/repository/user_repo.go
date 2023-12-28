package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_project/internal/constants"
	"go_project/internal/model/model"
	"go_project/internal/types"
	"go_project/internal/utils"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	UserRegister(ctx context.Context, req *types.UserRegisterReq) (*types.UserRegisterRes, error)                      // UserRegister 用户注册
	UserLogin(ctx context.Context, req *types.UserLoginReq) (*types.UserLoginRes, error)                               // UserLogin 用户登录
	UserUpdatePassword(ctx context.Context, req *types.UserUpdatePasswordReq) (*types.UserUpdatePasswordRes, error)    // UserUpdatePassword 用户修改密码
	AdminDeleteUserById(ctx context.Context, req *types.AdminDeleteUserByIdReq) (*types.AdminDeleteUserByIdRes, error) // AdminDeleteUserById 管理员删除用户
	GetUserList(ctx context.Context, req *types.GetUserListReq) (*types.GetUserListRes, error)                         // GetUserList 获取用户列表
}

type UserRepositoryImpl struct {
	UserDB    *gorm.DB
	UserRedis *redis.Client
}

// NewUserRepository 初始化 CourseRepositoryImpl 结构体实例的函数
func NewUserRepository(userDB *gorm.DB, userRedis *redis.Client) UserRepository {
	return &UserRepositoryImpl{
		UserDB:    userDB,
		UserRedis: userRedis,
	}
}

// UserRegister
// @Description 用户注册
// @Author Oberl-Fitzgerald 2023-12-27 09:38:24
// @Param  ctx context.Context
// @Param  req *types.UserRegisterReq
// @Return *types.UserRegisterRes
// @Return error
func (ur *UserRepositoryImpl) UserRegister(ctx context.Context, req *types.UserRegisterReq) (*types.UserRegisterRes, error) {
	// 判断传入参数是否为空,如果没传入type则默认为学生
	if req.Username == "" || req.Password == "" || req.Email == "" || req.Phone == "" {
		return &types.UserRegisterRes{
			Code: constants.ParamError.Code,
			Msg:  constants.ParamError.Message,
		}, nil
	}
	// 判断用户名是否已经存在
	var user model.User
	if findUserErr := ur.UserDB.WithContext(ctx).Where("phone = ?", req.Phone).First(&user).Error; findUserErr == nil {
		return &types.UserRegisterRes{
			Code: constants.UserExist.Code,
			Msg:  constants.UserExist.Message,
		}, nil
	}
	// 加密密码
	encodePWD, encodePWDErr := utils.EncodePassword(req.Password)
	if encodePWDErr != nil {
		return &types.UserRegisterRes{
			Code: constants.EncodePasswordError.Code,
			Msg:  constants.EncodePasswordError.Message,
		}, nil
	}
	// 创建用户
	newUser := model.User{
		UserName:  req.Username,
		Password:  encodePWD,
		Email:     req.Email,
		Phone:     req.Phone,
		RoleType:  int64(req.RoleType),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	if createUserErr := ur.UserDB.WithContext(ctx).Create(&newUser).Error; createUserErr != nil {
		return &types.UserRegisterRes{
			Code: constants.UserRegisterError.Code,
			Msg:  constants.UserRegisterError.Message,
		}, nil
	}
	return &types.UserRegisterRes{
		Code: constants.UserRegisterSuccess.Code,
		Msg:  constants.UserRegisterSuccess.Message,
		Data: &types.UserRegisterData{
			UserId: newUser.UserID,
		},
	}, nil

}

// UserLogin
// @Description 用户登录
// @Author Oberl-Fitzgerald 2023-12-27 11:23:59
// @Param  ctx context.Context
// @Param  req *types.UserLoginReq
// @Return *types.UserLoginRes
// @Return error
func (ur *UserRepositoryImpl) UserLogin(ctx context.Context, req *types.UserLoginReq) (*types.UserLoginRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ur.UserDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.UserLoginRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.UserLoginRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 获取用户角色
	roleName, _ := model.GetRoleType(user.RoleType)
	// 返回用户信息
	return &types.UserLoginRes{
		Code: constants.UserLoginSuccess.Code,
		Msg:  constants.UserLoginSuccess.Message,
		Data: &types.UserData{
			UserId:   user.UserID,
			Username: user.UserName,
			Email:    user.Email,
			Phone:    user.Phone,
			Role: types.RoleData{
				RoleName: roleName,
				RoleType: int(user.RoleType),
			},
		},
	}, nil
}

// UserUpdatePassword
// @Description 用户修改密码
// @Author Oberl-Fitzgerald 2023-12-27 11:48:31
// @Param  ctx context.Context
// @Param  req *types.UserUpdatePasswordReq
// @Return *types.UserUpdatePasswordRes
// @Return error
func (ur *UserRepositoryImpl) UserUpdatePassword(ctx context.Context, req *types.UserUpdatePasswordReq) (*types.UserUpdatePasswordRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ur.UserDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.UserUpdatePasswordRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.UserUpdatePasswordRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 加密新密码
	encodePWD, encodePWDErr := utils.EncodePassword(req.NewPassword)
	if encodePWDErr != nil {
		return &types.UserUpdatePasswordRes{
			Code: constants.EncodePasswordError.Code,
			Msg:  constants.EncodePasswordError.Message,
		}, nil
	}
	// 更新密码
	if updateUserErr := ur.UserDB.WithContext(ctx).Model(&user).Update("password", encodePWD).Error; updateUserErr != nil {
		return &types.UserUpdatePasswordRes{
			Code: constants.UserUpdatePassErr.Code,
			Msg:  constants.UserUpdatePassErr.Message,
		}, nil
	}
	// 获取用户角色
	roleName, _ := model.GetRoleType(user.RoleType)
	return &types.UserUpdatePasswordRes{
		Code: constants.UserUpdatePassSuc.Code,
		Msg:  constants.UserUpdatePassSuc.Message,
		Data: &types.UserData{
			UserId:   user.UserID,
			Username: user.UserName,
			Email:    user.Email,
			Phone:    user.Phone,
			Role: types.RoleData{
				RoleName: roleName,
				RoleType: int(user.RoleType),
			},
		},
	}, nil

}

// AdminDeleteUserById
// @Description 管理员删除用户 管理员可以删除管理员和学生
// @Author Oberl-Fitzgerald 2023-12-27 14:25:01
// @Param  ctx context.Context
// @Param  req *types.AdminDeleteUserByIdReq
// @Return *types.AdminDeleteUserByIdRes
// @Return error
func (ur *UserRepositoryImpl) AdminDeleteUserById(ctx context.Context, req *types.AdminDeleteUserByIdReq) (*types.AdminDeleteUserByIdRes, error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := ur.UserDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.AdminDeleteUserByIdRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.AdminDeleteUserByIdRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 校验身份
	if user.RoleType != constants.AdminRole {
		return &types.AdminDeleteUserByIdRes{
			Code: constants.UserNotAdmin.Code,
			Msg:  constants.UserNotAdmin.Message,
		}, nil
	}
	// 根据id查询对应用户
	var userById model.User
	if findUserErr := ur.UserDB.WithContext(ctx).Where("user_id = ? and status = ? ", req.UserId, constants.StatusNormal).First(&userById).Error; findUserErr != nil {
		return &types.AdminDeleteUserByIdRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 删除用户
	if deleteUserErr := ur.UserDB.WithContext(ctx).Model(&userById).Update("status", constants.StatusBan).Error; deleteUserErr != nil {
		return &types.AdminDeleteUserByIdRes{
			Code: constants.UserDeleteErr.Code,
			Msg:  constants.UserDeleteErr.Message,
		}, nil
	}
	return &types.AdminDeleteUserByIdRes{
		Code: constants.UserDeleteSuc.Code,
		Msg:  constants.UserDeleteSuc.Message,
	}, nil
}

// GetUserList
// @Description 获取用户列表,分页查询，支持按照用户名查询
// @Author Oberl-Fitzgerald 2023-12-27 14:43:28
// @Param  ctx context.Context
// @Param  req *types.GetUserListReq
// @Return *types.GetUserListRes
// @Return error
func (ur *UserRepositoryImpl) GetUserList(ctx context.Context, req *types.GetUserListReq) (*types.GetUserListRes, error) {
	var userList []model.User
	var total int64
	// 计算偏移量
	offset := (req.Page - 1) * req.Size
	// 查询用户列表
	if FindUserListErr := ur.UserDB.WithContext(ctx).Where("status = ? and user_name like ?", constants.StatusNormal, "%"+req.Keyword+"%").Offset(offset).Limit(req.Size).Find(&userList).Error; FindUserListErr != nil {
		return &types.GetUserListRes{
			Code: constants.UserListErr.Code,
			Msg:  constants.UserListErr.Message,
		}, nil
	}
	// 查询用户总数
	if FindUserTotalErr := ur.UserDB.WithContext(ctx).Model(&model.User{}).Where("status = ? and user_name like ?", constants.StatusNormal, "%"+req.Keyword+"%").Count(&total).Error; FindUserTotalErr != nil {
		return &types.GetUserListRes{
			Code: constants.UserCountErr.Code,
			Msg:  constants.UserCountErr.Message,
		}, nil
	}
	// 将用户列表转换为返回数据
	var userListData []*types.UserData
	for i := 0; i < len(userList); i++ {
		roleName, _ := model.GetRoleType(userList[i].RoleType)
		userListData = append(userListData, &types.UserData{
			UserId:   userList[i].UserID,
			Username: userList[i].UserName,
			Email:    userList[i].Email,
			Phone:    userList[i].Phone,
			Role: types.RoleData{
				RoleName: roleName,
				RoleType: int(userList[i].RoleType),
			},
		})
	}
	return &types.GetUserListRes{
		Code:  constants.UserListSuc.Code,
		Msg:   constants.UserListSuc.Message,
		Data:  userListData,
		Total: int(total),
	}, nil
}
