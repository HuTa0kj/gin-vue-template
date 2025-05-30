package global

const (
	// 正常请求（2）
	CodeSuccessMsg = "Successful"

	// 字段异常（3）
	CodeParameterMissingMsg    = "Parameter missing"
	CodeInformationNotFoundMsg = "Information not found"
	CodeParameterIllegalMsg    = "Parameter illegal"

	// 认证错误 (4)
	CodeUnauthorizedMsg        = "Unauthorized"
	CodeLoginFailMsg           = "Login failed"
	CodeTokenParsingFailedMsg  = "Token parsing failed"
	CodeTokenInvalidMsg        = "Token invalid"
	CodeTokenExpiredMsg        = "Token expired"
	CodeOriginPasswordErrorMsg = "Origin password error"
	CodeUserDisableMsg         = "User disabled"

	// 业务错误 (5)
	CodeDuplicateUsernameMsg = "Username already exists"
	CodeInviteCodeInvalidMsg = "Invite code invalid"
	CodeInviteRoleErrorMsg   = "Invite role error"

	// 数据库错误（6）
	CodeDatabaseInsertErrorMsg = "Data insert error"
	CodeDatabaseUpdateErrorMsg = "Data update error"
	CodeDatabaseSelectErrorMsg = "Data select error"
)
