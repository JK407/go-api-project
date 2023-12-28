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

type OrderRepository interface {
	PlaceOrder(ctx context.Context, req *types.PlaceOrderReq) (resp *types.PlaceOrderRes, err error)       // PlaceOrder 下单
	GetOrderList(ctx context.Context, req *types.GetOrderListReq) (resp *types.GetOrderListRes, err error) // GetOrderList 获取订单列表
	CancelOrder(ctx context.Context, req *types.CancelOrderReq) (resp *types.CancelOrderRes, err error)    // CancelOrder 取消订单
	OrderInfo(ctx context.Context, req *types.OrderInfoReq) (resp *types.OrderInfoRes, err error)          // OrderInfo 订单详情
}

type OrderRepositoryImpl struct {
	OrderDB    *gorm.DB
	OrderRedis *redis.Client
}

// OrderInfo
// @Description 订单详情
// @Author Oberl-Fitzgerald 2023-12-28 17:38:51
// @Param  ctx context.Context
// @Param  req *types.OrderInfoReq
// @Return resp
// @Return err
func (o OrderRepositoryImpl) OrderInfo(ctx context.Context, req *types.OrderInfoReq) (resp *types.OrderInfoRes, err error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := o.OrderDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.OrderInfoRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.OrderInfoRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 查询订单是否存在
	var order model.Order
	if err = o.OrderDB.WithContext(ctx).Where("order_id = ? and user_id = ? and status = ? ", req.OrderId, user.UserID, constants.StatusNormal).First(&order).Error; err != nil {
		return &types.OrderInfoRes{
			Code: constants.OrderNotExist.Code,
			Msg:  constants.OrderNotExist.Message,
		}, nil
	}
	// 查询课程详情
	var course model.Course
	if err = o.OrderDB.WithContext(ctx).Where("course_id = ? and status = ? ", order.CourseID, constants.StatusNormal).First(&course).Error; err != nil {
		return &types.OrderInfoRes{
			Code: constants.CourseNotExist.Code,
			Msg:  constants.CourseNotExist.Message,
		}, nil
	}
	// 封装返回数据
	orderData := &types.OrderData{
		OrderId:       order.OrderID,
		UserId:        user.UserID,
		CourseId:      order.CourseID,
		TotalPrice:    order.TotalPrice,
		PurchaseCount: order.PurchaseCount,
	}
	courseData := &types.CourseData{
		CourseId:    course.CourseID,
		CourseName:  course.CourseName,
		Description: course.Description,
		Price:       course.Price,
		UserId:      course.UserID,
	}
	// 获取用户角色
	roleName, _ := model.GetRoleType(user.RoleType)
	userData := &types.UserData{
		UserId:   user.UserID,
		Username: user.UserName,
		Email:    user.Email,
		Phone:    user.Phone,
		Role: types.RoleData{
			RoleName: roleName,
			RoleType: int(user.RoleType),
		},
	}
	return &types.OrderInfoRes{
		Code: constants.OrderInfoSuc.Code,
		Msg:  constants.OrderInfoSuc.Message,
		Data: &types.OrderInfoData{
			OrderData:  orderData,
			CourseData: courseData,
			UserData:   userData,
		},
	}, nil
}

// CancelOrder
// @Description 取消订单
// @Author Oberl-Fitzgerald 2023-12-28 17:29:59
// @Param  ctx context.Context
// @Param  req *types.CancelOrderReq
// @Return resp
// @Return err
func (o OrderRepositoryImpl) CancelOrder(ctx context.Context, req *types.CancelOrderReq) (resp *types.CancelOrderRes, err error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := o.OrderDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.CancelOrderRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.CancelOrderRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 查询订单是否存在
	var order model.Order
	if err = o.OrderDB.WithContext(ctx).Where("order_id = ? and user_id = ? and status = ? ", req.OrderId, user.UserID, constants.StatusNormal).First(&order).Error; err != nil {
		return &types.CancelOrderRes{
			Code: constants.OrderNotExist.Code,
			Msg:  constants.OrderNotExist.Message,
		}, nil
	}
	// 修改订单状态为已取消
	if err = o.OrderDB.WithContext(ctx).Model(&order).Where("order_id = ? and user_id = ? and status = ? ", req.OrderId, user.UserID, constants.StatusNormal).Update("status", constants.StatusBan).Error; err != nil {
		return &types.CancelOrderRes{
			Code: constants.OrderCancelErr.Code,
			Msg:  constants.OrderCancelErr.Message,
		}, nil
	}
	return &types.CancelOrderRes{
		Code: constants.OrderCancelSuc.Code,
		Msg:  constants.OrderCancelSuc.Message,
	}, nil
}

// GetOrderList
// @Description 获取订单列表
// @Author Oberl-Fitzgerald 2023-12-28 17:24:40
// @Param  ctx context.Context
// @Param  req *types.GetOrderListReq
// @Return resp
// @Return err
func (o OrderRepositoryImpl) GetOrderList(ctx context.Context, req *types.GetOrderListReq) (resp *types.GetOrderListRes, err error) {
	var orderList []*model.Order
	var total int64
	// 计算偏移量
	offset := (req.Page - 1) * req.Size
	// 查询订单列表
	if err = o.OrderDB.WithContext(ctx).Where("status = ? ", constants.StatusNormal).Offset(offset).Limit(req.Size).Find(&orderList).Error; err != nil {
		return &types.GetOrderListRes{
			Code: constants.OrderListErr.Code,
			Msg:  constants.OrderListErr.Message,
		}, nil
	}
	// 查询订单总数
	if err = o.OrderDB.WithContext(ctx).Model(&model.Order{}).Where("status = ? ", constants.StatusNormal).Count(&total).Error; err != nil {
		return &types.GetOrderListRes{
			Code: constants.OrderCountErr.Code,
			Msg:  constants.OrderCountErr.Message,
		}, nil
	}
	// 封装返回数据
	var orderDataList []*types.OrderData
	for _, order := range orderList {
		orderDataList = append(orderDataList, &types.OrderData{
			OrderId:       order.OrderID,
			UserId:        order.UserID,
			CourseId:      order.CourseID,
			TotalPrice:    order.TotalPrice,
			PurchaseCount: order.PurchaseCount,
		})
	}
	return &types.GetOrderListRes{
		Code:  constants.OrderListSuc.Code,
		Msg:   constants.OrderListSuc.Message,
		Data:  orderDataList,
		Total: int(total),
	}, nil

}

// PlaceOrder
// @Description 下单
// @Author Oberl-Fitzgerald 2023-12-28 16:43:51
// @Param  ctx context.Context
// @Param  req *types.PlaceOrderReq
// @Return resp
// @Return err
func (o OrderRepositoryImpl) PlaceOrder(ctx context.Context, req *types.PlaceOrderReq) (resp *types.PlaceOrderRes, err error) {
	// 根据手机号查询用户
	var user model.User
	if findUserErr := o.OrderDB.WithContext(ctx).Where("phone = ? and status = ? ", req.Phone, constants.StatusNormal).First(&user).Error; findUserErr != nil {
		return &types.PlaceOrderRes{
			Code: constants.UserNotExist.Code,
			Msg:  constants.UserNotExist.Message,
		}, nil
	}
	// 校验密码
	if !utils.ValidatePassword(user.Password, req.Password) {
		return &types.PlaceOrderRes{
			Code: constants.PasswordError.Code,
			Msg:  constants.PasswordError.Message,
		}, nil
	}
	// 查询课程是否存在
	var course model.Course
	if err = o.OrderDB.WithContext(ctx).Where("course_id = ? and status = ? ", req.CourseId, constants.StatusNormal).First(&course).Error; err != nil {
		return &types.PlaceOrderRes{
			Code: constants.CourseNotExist.Code,
			Msg:  constants.CourseNotExist.Message,
		}, nil
	}
	// 计算订单价格
	orderPrice := course.Price * float64(req.PurchaseCount)
	// 创建订单
	order := model.Order{
		UserID:        user.UserID,
		CourseID:      req.CourseId,
		TotalPrice:    orderPrice,
		PurchaseCount: req.PurchaseCount,
	}
	if err = o.OrderDB.WithContext(ctx).Create(&order).Error; err != nil {
		return &types.PlaceOrderRes{
			Code: constants.PlaceOrderErr.Code,
			Msg:  constants.PlaceOrderErr.Message,
		}, nil
	}
	return &types.PlaceOrderRes{
		Code: constants.PlaceOrderSuc.Code,
		Msg:  constants.PlaceOrderSuc.Message,
		Data: &types.OrderData{
			OrderId:       order.OrderID,
			UserId:        user.UserID,
			CourseId:      order.CourseID,
			TotalPrice:    order.TotalPrice,
			PurchaseCount: order.PurchaseCount,
		},
	}, nil
}

func NewOrderRepository(db *gorm.DB, redis *redis.Client) OrderRepository {
	return &OrderRepositoryImpl{
		OrderDB:    db,
		OrderRedis: redis,
	}
}
