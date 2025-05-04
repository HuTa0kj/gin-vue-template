package global

const (
	// 正常请求（2）
	CodeSuccess    = 2000
	CodeSuccessMsg = "Successful"

	// 字段异常（3）
	CodeParameterMissing       = 3001 // 字段缺失
	CodeParameterMissingMsg    = "Parameter missing"
	CodeInformationNotFound    = 3002 // 信息不存在
	CodeInformationNotFoundMsg = "Information not found"
	CodeParameterIllegal       = 3003 // 参数不合法
	CodeParameterIllegalMsg    = "Parameter illegal"

	// 认证错误 (4)
	CodeUnauthorized           = 4001 // 未授权
	CodeUnauthorizedMsg        = "Unauthorized"
	CodeLoginFail              = 4002 // 登录失败（账号密码错误）
	CodeLoginFailMsg           = "Login failed"
	CodeTokenParsingFailed     = 4003 // Token 解析失败
	CodeTokenParsingFailedMsg  = "Token parsing failed"
	CodeTokenInvalid           = 4004 // Token 无效
	CodeTokenInvalidMsg        = "Token invalid"
	CodeTokenExpired           = 4005 // Token 过期
	CodeTokenExpiredMsg        = "Token expired"
	CodeOriginPasswordError    = 4006 // 原始密码错误
	CodeOriginPasswordErrorMsg = "Origin password error"
	CodeUserDisable            = 4007 // 用户被禁用
	CodeUserDisableMsg         = "User disabled"

	// 业务错误 (5)
	CodeDuplicateUsername    = 5001
	CodeDuplicateUsernameMsg = "Username already exists"
	CodeInviteCodeInvalid    = 5002
	CodeInviteCodeInvalidMsg = "Invite code invalid"
	CodeInviteRoleError      = 5003
	CodeInviteRoleErrorMsg   = "Invite role error"

	// 数据库错误（6）
	CodeDatabaseInsertError    = 6001
	CodeDatabaseInsertErrorMsg = "Data insert error"
	CodeDatabaseUpdateError    = 6002
	CodeDatabaseUpdateErrorMsg = "Data update error"
	CodeDatabaseSelectError    = 6003
	CodeDatabaseSelectErrorMsg = "Data select error"
)
