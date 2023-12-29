# go-api-project 技术考研学习平台

## 运行环境
- go 1.20 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" width="30" valign="middle" />
- mysql 8.0 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mysql/mysql-original-wordmark.svg" width="30" valign="middle" />
- redis <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/redis/redis-original-wordmark.svg" width="30" valign="middle" />

---


## 运行命令
```shell
go mod tidy
```
```shell
go run ./service.go
```
---

## 模块划分

### 模块一 用户模块
- [x] UserRegister // 用户注册 
> 用户注册接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
// POST
curl --location '127.0.0.1:8888/users/register' \ 
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"测试", //用户名
    "password":"123456", //用户密码
    "email":"111111@qq.com", //用户邮箱
    "phone":"4", //用户手机号
    "role_type":1 //角色类型:1学生，2管理员
}'
```
- [x] UserLogin // 用户登录 (手机号+密码)
> 用户登录接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/users/login' \
--header 'Content-Type: application/json' \
--data '{
    "password": "123456", //密码
    "phone": "3" //手机号
}'
```
- [x] UserUpdatePassword // 用户自己修改密码
> 修改密码接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/users/updatePassword' \
--header 'Content-Type: application/json' \
--data '{
    "password": "123456", //密码
    "phone": "1111112", //手机号
    "new_password":"1234" //新密码
}'
```
- [x] AdminDeleteUserById // 管理员可以通过用户id删除用户 (假删除)
> 管理员注销用户接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/deleteUserById' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":3, //用户id
    "phone":"1", //用户手机号
    "password":"123456" //用户密码
}'
```
- [x] GetUserList // 查询用户列表 (分页查询，关键字<font face="隶书" color=##FF0000 size=3>**用户名**</font>查询)
> 查询用户列表接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/getUserList' \
--header 'Content-Type: application/json' \
--data '{
    "page":1, //页数
    "size":5, //每页数量
    "keyword":"" //关键词
}'
```

### 模块二 课程模块
- [x] AdminAddCourse 管理员添加课程 (课程名称可以重复)
> 管理员添加课程接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/addCourse' \
--header 'Content-Type: application/json' \
--data '{
    	"course_name":"C课程", //课程名称
		"description":"测试课程", //课程描述
		"price":62.3, //课程价格
		"phone":"1", //用户手机号
		"password":"123456" //用户密码
}'
```
- [x] AdminUpdateCourse 管理员修改课程价格
> 管理员修改课程价格接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/updateCourse' \
--header 'Content-Type: application/json' \
--data '{
		"price":9.12, // 课程价格
		"phone":"4", // 用户手机号
		"password":"123456", // 用户密码
        "course_id":1 //课程id
}'
```
- [x] AdminDelCourse 管理员删除课程 (同时清除绑定表中的对应数据)
> 管理员删除课程接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/delCourse' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"4", //用户手机号
		"password":"123456", //用户密码
        "course_id":2 //课程id
}'
```
- [x] GetCourseList 获取课程列表 (分页查询，关键字[课程名]查询)
> 获取课程列表接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/getCourseList' \
--header 'Content-Type: application/json' \
--data '{
    "page":1, //页数
    "size":3,// 每页数量
    "keyword":"课程"  // 关键词
}'
```

### 模块三 课程类别模块
- [x] AdminAddCategory 管理员添加分类 (课程类别名称不能重复)
> 管理员添加分类接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/addCategory' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"1", //用户手机号
		"password":"123456", //用户密码
        "category_name":"Node栏目" //栏目名称
}'
```
- [x] AdminUpdateCategory 管理员修改分类名称
> 管理员修改分类名称接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/updateCategory' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"4", // 用户手机号
		"password":"123456", // 用户密码
        "category_id":1, // 栏目id
        "category_name":"Node1栏目" // 新栏目名称
}'
```
- [x] AdminBindCourseCategory 管理员绑定课程分类
> 管理员绑定课程分类接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/bindCourseCategory' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"4", // 手机号
		"password":"123456", // 密码
        "category_id":2, // 类别id
        "course_id":3 //课程id
}'
```
- [x] AdminDelCategory 管理员删除分类 (同时清除绑定表中的对应数据)
> 管理员删除分类接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/delCategory' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"4", //手机号
		"password":"123456", //密码
        "category_id":10 //栏目id
}'
```
- [x] GetCategoryList 获取分类列表 (分页查询，关键字[课程类别名]查询)
> 获取分类列表接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/admin/getCategoryList' \
--header 'Content-Type: application/json' \
--data '{
    "page":1, // 页数
    "size":5, // 每页数据数
    "keyword":"" // 关键词
}'
```

### 模块四 订单模块
- [x] PlaceOrder 下单
> 下单接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/order/placeOrder' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"3", // 手机号
		"password":"123456", // 密码
        "purchase_count":2, // 购买数量
        "course_id":3 // 购买课程id
}'
```
- [x] GetOrderList 获取订单列表 (分页列表)
> 获取订单列表接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/order/getOrderList' \
--header 'Content-Type: application/json' \
--data '{
    "page":1, // 页数
    "size":1 // 每页数据数
}'
```
- [x] CancelOrder 取消订单
> 取消订单接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/order/cancelOrder' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"3", // 手机号
		"password":"123456", // 密码
        "order_id":1 // 订单id
}'
```
- [x] OrderInfo 订单详情
> 订单详情接口 <font face="隶书" color=#be905c size=3>**POST**</font>
```shell
curl --location '127.0.0.1:8888/order/orderInfo' \
--header 'Content-Type: application/json' \
--data '{
		"phone":"3", // 手机号
		"password":"123456", // 密码
        "order_id":2 // 订单id
}'
```

---

## 联系我

如果您有任何问题或建议，请随时与我联系。

- GitHub: [JK407](https://github.com/JK407) <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original-wordmark.svg" width="30" valign="middle" />
- Email: 1467511770@qq.com




          
          
          