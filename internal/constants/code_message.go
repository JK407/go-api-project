package constants

type CodeMessage struct {
	Code    int
	Message string
	Comment string
}

var (
	Success = CodeMessage{Code: 200, Message: "success", Comment: "成功"}

	ParamError = CodeMessage{Code: 400, Message: "param error", Comment: "参数错误"}

	ServerError = CodeMessage{Code: 500, Message: "server error", Comment: "服务错误"}

	UserCreateError = CodeMessage{Code: 1000, Message: "user create error", Comment: "用户创建失败"}
	UserExist       = CodeMessage{Code: 1001, Message: "user exist", Comment: "用户已存在"}
)
