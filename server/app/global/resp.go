package global

const (
	// 正常请求（2）
	CodeSuccess    = 2000
	CodeSuccessMsg = "Successful"

	// 字段缺失（3）
	CodeParameterMissing    = 3001 // 字段缺失
	CodeParameterMissingMsg = "Parameter missing"

	// 认证错误 (4)
	CodeUnauthorized          = 4001 // 未授权
	CodeUnauthorizedMsg       = "Unauthorized"
	CodeTokenInvalid          = 4002 // Token 无效
	CodeTokenInvalidMsg       = "Token invalid"
	CodeTokenExpired          = 4003 // Token 过期
	CodeTokenExpiredMsg       = "Token expired"
	CodeTokenParsingFailed    = 4004 // Token 解析失败
	CodeTokenParsingFailedMsg = "Token parsing failed"
	CodeLoginFail             = 4005 // 登录失败（账号密码错误）
	CodeLoginFailMsg          = "Login failed"

	// 业务错误 (5)
	CodeInformationNotFound    = 5001 // 信息不存在
	CodeInformationNotFoundMsg = "Information not found"
)
