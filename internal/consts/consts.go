package consts

const (
	ProjectName              = "Go开源电商实战项目"         // 项目名称
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GTokenAdminPrefix        = "Admin:"             //gtoken登录 管理后台 前缀区分
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	CacheModeRedis           = 2
	BackendServerName        = "开源电商服系统"
	MultiLogin               = true
	ErrorLoginFailMsg        = "登录失败，账号或密码错误."
	GTokenExpireIn           = 10 * 24 * 60 * 60
)
