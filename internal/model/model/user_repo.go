package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_project/internal/constants"
	"go_project/internal/types"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	UserRegister(ctx context.Context, req *types.UserRegisterReq) (*types.UserRegisterRes, error) // UserRegister 用户注册
}

type UserRepositoryImpl struct {
	UserDB    *gorm.DB
	UserRedis *redis.Client
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
	var user User
	if err := ur.UserDB.WithContext(ctx).Where("phone = ?", req.Phone).First(&user).Error; err == nil {
		return &types.UserRegisterRes{
			Code: constants.UserExist.Code,
			Msg:  constants.UserExist.Message,
		}, nil
	}
	// 创建用户
	newUser := User{
		UserName:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Phone:     req.Phone,
		RoleType:  int64(req.RoleType),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	if err := ur.UserDB.WithContext(ctx).Create(&newUser).Error; err != nil {
		return &types.UserRegisterRes{
			Code: constants.UserCreateError.Code,
			Msg:  constants.UserCreateError.Message,
		}, nil
	}
	return &types.UserRegisterRes{
		Code: constants.Success.Code,
		Msg:  constants.Success.Message,
		Data: types.UserRegisterData{
			UserId: newUser.UserID,
		},
	}, nil

}
