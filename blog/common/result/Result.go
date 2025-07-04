package result

var (

	// OK
	OK  = response(200, "ok") // 通用成功
	Err = response(500, "")   // 通用错误

	// 服务级错误码
	ErrParam     = response(10001, "参数有误")
	ErrSignParam = response(10002, "签名参数有误")

	// 模块级错误码 - 用户模块
	ErrUserService     = response(20100, "用户服务异常")
	ErrUserNameExists  = response(20101, "用户名称已存在")
	ErrUserEmailExists = response(20102, "用户邮箱已存在")
	ErrUserCreate      = response(20103, "用户注册失败")
	ErrUserQuery       = response(20104, "用户查询失败")
	ErrUserOrPwd       = response(20105, "用户名或密码错误")

	// 库存模块
	ErrOrderService = response(20200, "订单服务异常")
	ErrOrderOutTime = response(20201, "订单超时")
)
