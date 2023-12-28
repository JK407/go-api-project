package constants

type CodeMessage struct {
	Code     int
	Message  string
	Annotate string
}

var (
	Success               = CodeMessage{Code: 200, Message: "success", Annotate: "成功"}
	UserRegisterSuccess   = CodeMessage{Code: 2001, Message: "user register success", Annotate: "用户注册成功"}
	UserLoginSuccess      = CodeMessage{Code: 2002, Message: "user login success", Annotate: "用户登录成功"}
	UserUpdatePassSuc     = CodeMessage{Code: 2003, Message: "user update password success", Annotate: "用户修改密码成功"}
	UserDeleteSuc         = CodeMessage{Code: 2004, Message: "user delete success", Annotate: "用户删除成功"}
	UserListSuc           = CodeMessage{Code: 2005, Message: "user list success", Annotate: "用户列表获取成功"}
	CourseAddSuc          = CodeMessage{Code: 2006, Message: "course add success", Annotate: "课程添加成功"}
	CourseUpdateSuc       = CodeMessage{Code: 2007, Message: "course update success", Annotate: "课程更新成功"}
	CourseDeleteSuc       = CodeMessage{Code: 2008, Message: "course delete success", Annotate: "课程删除成功"}
	CourseListSuc         = CodeMessage{Code: 2009, Message: "course list success", Annotate: "课程列表获取成功"}
	CategoryCreateSuc     = CodeMessage{Code: 2010, Message: "category create success", Annotate: "分类创建成功"}
	CourseCategoryBindSuc = CodeMessage{Code: 2011, Message: "course category bind success", Annotate: "课程分类绑定成功"}
	CategoryDeleteSuc     = CodeMessage{Code: 2012, Message: "category delete success", Annotate: "分类删除成功"}
	CategoryListSuc       = CodeMessage{Code: 2013, Message: "category list success", Annotate: "分类列表获取成功"}
	PlaceOrderSuc         = CodeMessage{Code: 2014, Message: "place order success", Annotate: "下单成功"}
	OrderListSuc          = CodeMessage{Code: 2015, Message: "order list success", Annotate: "订单列表获取成功"}
	OrderCancelSuc        = CodeMessage{Code: 2016, Message: "order cancel success", Annotate: "订单取消成功"}
	OrderInfoSuc          = CodeMessage{Code: 2017, Message: "order info success", Annotate: "订单详情获取成功"}

	ParamError              = CodeMessage{Code: 400, Message: "param error", Annotate: "参数错误"}
	UserRegisterError       = CodeMessage{Code: 4001, Message: "user register error", Annotate: "用户注册失败"}
	UserLoginError          = CodeMessage{Code: 4002, Message: "user login error", Annotate: "用户登录失败"}
	UserUpdatePassErr       = CodeMessage{Code: 4003, Message: "user update password error", Annotate: "用户修改密码失败"}
	UserNotAdmin            = CodeMessage{Code: 4004, Message: "user not admin", Annotate: "用户不是管理员"}
	UserDeleteErr           = CodeMessage{Code: 4005, Message: "user delete error", Annotate: "用户删除失败"}
	UserListErr             = CodeMessage{Code: 4006, Message: "user list error", Annotate: "用户列表获取失败"}
	UserCountErr            = CodeMessage{Code: 4007, Message: "user count error", Annotate: "用户总数获取失败"}
	CourseAddErr            = CodeMessage{Code: 4008, Message: "course add error", Annotate: "课程添加失败"}
	CourseUpdateErr         = CodeMessage{Code: 4009, Message: "course update error", Annotate: "课程更新失败"}
	CourseNotExist          = CodeMessage{Code: 4010, Message: "course not exist", Annotate: "课程不存在"}
	CourseDeleteErr         = CodeMessage{Code: 4011, Message: "course delete error", Annotate: "课程删除失败"}
	CourseListErr           = CodeMessage{Code: 4012, Message: "course list error", Annotate: "课程列表获取失败"}
	CourseCountErr          = CodeMessage{Code: 4013, Message: "course count error", Annotate: "课程总数获取失败"}
	CategoryCreateErr       = CodeMessage{Code: 4014, Message: "category create error", Annotate: "分类创建失败"}
	CategoryNameExist       = CodeMessage{Code: 4015, Message: "category name exist", Annotate: "分类名称已存在"}
	CategoryNotExist        = CodeMessage{Code: 4016, Message: "category not exist", Annotate: "分类不存在"}
	CourseCategoryBindErr   = CodeMessage{Code: 4017, Message: "course category bind error", Annotate: "课程分类绑定失败"}
	CourseCategoryBindExist = CodeMessage{Code: 4018, Message: "course category bind exist", Annotate: "课程分类绑定已存在"}
	CategoryDeleteErr       = CodeMessage{Code: 4019, Message: "category delete error", Annotate: "分类删除失败"}
	CategoryListErr         = CodeMessage{Code: 4020, Message: "category list error", Annotate: "分类列表获取失败"}
	CategoryCountErr        = CodeMessage{Code: 4021, Message: "category count error", Annotate: "分类总数获取失败"}
	PlaceOrderErr           = CodeMessage{Code: 4022, Message: "place order error", Annotate: "下单失败"}
	OrderListErr            = CodeMessage{Code: 4023, Message: "order list error", Annotate: "订单列表获取失败"}
	OrderCountErr           = CodeMessage{Code: 4024, Message: "order count error", Annotate: "订单总数获取失败"}
	OrderNotExist           = CodeMessage{Code: 4025, Message: "order not exist", Annotate: "订单不存在"}
	OrderCancelErr          = CodeMessage{Code: 4026, Message: "order cancel error", Annotate: "订单取消失败"}

	ServerError = CodeMessage{Code: 500, Message: "server error", Annotate: "服务错误"}

	UserExist    = CodeMessage{Code: 1000, Message: "user exist", Annotate: "用户已存在"}
	UserNotExist = CodeMessage{Code: 1001, Message: "user not exist", Annotate: "用户不存在"}

	EncodePasswordError = CodeMessage{Code: 3000, Message: "encode password error", Annotate: "密码加密失败"}
	PasswordError       = CodeMessage{Code: 3001, Message: "password error", Annotate: "密码错误"}

	RoleTypeNotExist = CodeMessage{Code: 6000, Message: "role type not exist", Annotate: "角色类型不存在"}
)

const (
	StatusNormal = 0
	StatusBan    = 1
	StudentRole  = 1
	AdminRole    = 2
)
