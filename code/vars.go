package code

var (
	OK                        = add(0)   // 正确
	NotModifiedErrCode        = add(304) // 木有改动
	TemporaryRedirectErrCode  = add(307) // 撞车跳转
	RequestErrCode            = add(400) // 请求错误
	UnauthorizedErrCode       = add(401) // 未认证
	AccessDeniedErrCode       = add(403) // 访问权限不足
	NothingFoundErrCode       = add(404) // 啥都木有
	MethodNotAllowedErrCode   = add(405) // 不支持该方法
	ConflictErrCode           = add(409) // 冲突
	CanceledErrCode           = add(498) // 客户端取消请求
	ServerErrCode             = add(500) // 服务器错误
	ServiceUnavailableErrCode = add(503) // 过载保护,服务暂不可用
	DeadlineErrCode           = add(504) // 服务调用超时

)
