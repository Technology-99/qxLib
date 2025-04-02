package qxCodes

// note: qx内部开发引擎部分状态码 1000-5999
// note: 如果是业务上的错误码，请从10000,20000,30000,40000,50000开头

// note: 补充信息性状态码 1000～1999
// note: 补充成功状态码 2000～2999
// note: 补充重定向状态码 3000～3999
// note: 补充客户端错误  4000～4999
// note: 补充服务器错误 5000～5999

const (
	QxEngineStatusContinue                      int32 = 1000 // 语义: 继续。
	QxEngineStatusSwitchingProtocols            int32 = 1001 // 语义: 服务器同意客户段的协议升级请求
	QxEngineStatusProcessing                    int32 = 1002 // 语义: 服务器已接收请求，但是仍然在处理中
	QxEngineStatusEarlyHints                    int32 = 1003 // 语义: 服务器提前发送头部信息（用于优化请求）
	QxEngineStatusAccessTokenInvalid            int32 = 1004 // 语义: 错误的访问凭证
	QxEngineStatusAccessExpired                 int32 = 1005 // 语义: 访问凭证已过期
	QxEngineStatusRefreshTokenInvalid           int32 = 1006 // 语义: 错误的刷新凭证
	QxEngineStatusRefreshTokenExpired           int32 = 1007 // 语义: 刷新凭证已过期
	QxEngineStatusAccountNotFound               int32 = 1008 // 语义: 找不到该账号
	QxEngineStatusDisabled                      int32 = 1009 // 语义: 资源被禁用
	QxEngineStatusNotSetWhiteList               int32 = 1010 // 语义: 请联系系统管理员设置ip白名单
	QxEngineStatusNotWhiteList                  int32 = 1011 // 语义: ip %s, 不在白名单内
	QxEngineStatusNotFoundAuthorization         int32 = 1012 // 语义: 没找到验证头
	QxEngineStatusNotFoundXAccessKeyFor         int32 = 1013 // 语义: 没找到api密钥头
	QxEngineStatusNotFoundMetadata              int32 = 1014 // 语义: rpc通讯没找到元数据
	QxEngineStatusKeyChangeExpired              int32 = 1015 // 语义: 密钥交换已过期
	QxEngineStatusTokenNotActiveYet             int32 = 1016 // 语义: 凭证尚未激活
	QxEngineStatusSecurityTokenInvalid          int32 = 1017 // 语义: 安全令牌错误
	QxEngineStatusSecurityTokenExpired          int32 = 1018 // 语义: 安全令牌过期
	QxEngineStatusAccountOrPwdInvalid           int32 = 1019 // 语义: 账号或密码错误
	QxEngineStatusSasServiceNotBootstrap        int32 = 1020 // 语义: 存储中心尚未完成初始化
	QxEngineStatusOK                            int32 = 2000 // 语义: 请求成功。
	QxEngineStatusCreated                       int32 = 2001 // 语义: 已创建。
	QxEngineStatusAccepted                      int32 = 2002 // 语义: 已接受。
	QxEngineStatusNonAuthoritativeInfo          int32 = 2003 // 语义: 非授权信息。请求成功。但返回的meta信息不在原始的服务器，而是一个副本
	QxEngineStatusNoContent                     int32 = 2004 // 语义: 无内容。
	QxEngineStatusResetContents                 int32 = 2005 // 语义: 重置内容。
	QxEngineStatusPartialContent                int32 = 2006 // 语义: 部分内容。
	QxEngineStatusMultiStatus                   int32 = 2007 // 语义: WebDAV扩展，表示多个状态
	QxEngineStatusAlreadyReported               int32 = 2008 // 语义: WebDAV扩展，资源已存在
	QxEngineStatusIMUsed                        int32 = 2026 // 语义: 服务器完成请求，但返回的是变更后的资源
	QxEngineStatusNotTrustDevice                int32 = 2027 // 语义: 不是可信设备
	QxEngineStatusAccountExists                 int32 = 2028 // 语义: 账号已经存在
	QxEngineStatusResourceExists                int32 = 2029 // 语义: 资源已存在
	QxEngineStatusUnsupportedAlgorithmModel     int32 = 2030 // 语义: 不支持的算法模型
	QxEngineStatusUnsupportedSignMethod         int32 = 2031 // 语义: 不支持的签名算法
	QxEngineStatusMultipleChoices               int32 = 3000 // 语义: 多种选择。
	QxEngineStatusMovedPermanently              int32 = 3001 // 语义: 永久移动。
	QxEngineStatusFound                         int32 = 3002 // 语义: 临时移动。
	QxEngineStatusSeeOther                      int32 = 3003 // 语义: 查看其它地址。
	QxEngineStatusNotModified                   int32 = 3004 // 语义: 未修改。
	QxEngineStatusUseProxy                      int32 = 3005 // 语义: 使用代理。
	QxEngineUnused                              int32 = 3006 // 语义: 已废弃
	QxEngineStatusTemporaryRedirect             int32 = 3007 // 语义: 临时重定向。
	QxEngineStatusPermanentRedirect             int32 = 3008 // 语义: 永久重定向（类似301，但POST不会变GET）
	QxEngineStatusRedirectSuccess               int32 = 3009 // 语义: 重定向成功
	QxEngineStatusBadRequest                    int32 = 4000 // 语义: 客户端请求参数错误
	QxEngineStatusUnauthorized                  int32 = 4001 // 语义: 请求要求用户的身份认证
	QxEngineStatusPaymentRequired               int32 = 4002 // 语义: 保留，将来使用
	QxEngineStatusForbidden                     int32 = 4003 // 语义: 服务器理解请求客户端的请求，但是拒绝执行此请求
	QxEngineStatusNotFound                      int32 = 4004 // 语义: 服务器无法根据客户端的请求找到资源（网页）。
	QxEngineStatusMethodNotAllowed              int32 = 4005 // 语义: 客户端请求中的方法被禁止
	QxEngineStatusNotAcceptable                 int32 = 4006 // 语义: 服务器无法根据客户端请求的内容特性完成请求
	QxEngineStatusProxyAuthRequired             int32 = 4007 // 语义: 请求要求代理的身份认证，与401类似，但请求者应当使用代理进行授权
	QxEngineStatusRequestTimeout                int32 = 4008 // 语义: 服务器等待客户端发送的请求时间过长，超时
	QxEngineStatusConflict                      int32 = 4009 // 语义: 服务器完成客户端的 PUT 请求时可能返回此代码，服务器处理请求时发生了冲突
	QxEngineStatusGone                          int32 = 4010 // 语义: 客户端请求的资源已经不存在。
	QxEngineStatusLengthRequired                int32 = 4011 // 语义: 服务器无法处理客户端发送的不带Content-Length的请求信息
	QxEngineStatusPreconditionFailed            int32 = 4012 // 语义: 客户端请求信息的先决条件错误
	QxEngineStatusRequestEntityTooLarge         int32 = 4013 // 语义: 由于请求的实体过大，服务器无法处理，因此拒绝请求。为防止客户端的连续请求，服务器可能会关闭连接。如果只是服务器暂时无法处理，则会包含一个Retry-After的响应信息
	QxEngineStatusRequestURITooLong             int32 = 4014 // 语义: 请求的URI过长（URI通常为网址），服务器无法处理
	QxEngineStatusUnsupportedMediaType          int32 = 4015 // 语义: 服务器无法处理请求附带的媒体格式
	QxEngineStatusRequestedRangeNotSatisfiable  int32 = 4016 // 语义: 客户端请求的范围无效
	QxEngineStatusExpectationFailed             int32 = 4017 // 语义: 服务器无法满足请求头中 Expect 字段指定的预期行为。
	QxEngineStatusTeapot                        int32 = 4018 // 语义: 愚人节彩蛋
	QxEngineStatusMisdirectedRequest            int32 = 4021 // 语义: 服务器拒绝处理该请求
	QxEngineStatusUnprocessableEntity           int32 = 4022 // 语义: 请求格式正确，但数据有误
	QxEngineStatusLocked                        int32 = 4023 // 语义: 资源被锁定(如WebDAV资源锁)
	QxEngineStatusFailedDependency              int32 = 4024 // 语义: 由于前一个请求失败，当前请求也无法执行
	QxEngineStatusTooEarly                      int32 = 4025 // 语义: 服务器拒绝处理可能重放的请求
	QxEngineStatusUpgradeRequired               int32 = 4026 // 语义: 服务器要求客户端升级协议(如HTTP/1.1->HTTP/2)
	QxEngineStatusPreconditionRequired          int32 = 4028 // 语义: 需要先满足条件请求（如If-Match头）
	QxEngineStatusTooManyRequests               int32 = 4029 // 语义: 触发限流
	QxEngineStatusRequestHeaderFieldsTooLarge   int32 = 4031 // 语义: 请求头字段过大
	QxEngineStatusUnavailableForLegalReasons    int32 = 4051 // 语义: 由于法律原因，资源不可用
	QxEngineStatusAccountOrPasswordInvalid      int32 = 4052 // 语义: 账号或者密码错误
	QxEngineStatusPermissionDenied              int32 = 4053 // 语义: 权限不足
	QxEngineStatusNotStandardBase64Format       int32 = 4054 // 语义: 不是标准的base64格式
	QxEngineStatusKmsNotInit                    int32 = 4055 // 语义: 租户的kms尚未完成初始化
	QxEngineStatusFunctionUnavailable           int32 = 4056 // 语义: 功能不可用
	QxEngineStatusCodeInvalid                   int32 = 4057 // 语义: 验证码错误
	QxEngineStatusInternalServerError           int32 = 5000 // 语义: 服务器内部错误，无法完成请求
	QxEngineStatusNotImplemented                int32 = 5001 // 语义: 服务器不支持请求的功能，无法完成请求
	QxEngineStatusBadGateway                    int32 = 5002 // 语义: 作为网关或者代理工作的服务器尝试执行请求时，从远程服务器接收到了一个无效的响应
	QxEngineStatusServiceUnavailable            int32 = 5003 // 语义: 由于超载或系统维护，服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的Retry-After头信息中
	QxEngineStatusGatewayTimeout                int32 = 5004 // 语义: 充当网关或代理的服务器，未及时从远端服务器获取请求
	QxEngineStatusHTTPVersionNotSupported       int32 = 5005 // 语义: 服务器不支持请求的HTTP协议的版本，无法完成处理
	QxEngineStatusVariantAlsoNegotiates         int32 = 5006 // 语义: 服务器内部配置错误
	QxEngineStatusInsufficientStorage           int32 = 5007 // 语义: 服务器存储空间不足
	QxEngineStatusLoopDetected                  int32 = 5008 // 语义: 服务器监测到无限循环
	QxEngineStatusNotExtended                   int32 = 5010 // 语义: 服务器需要扩展请求
	QxEngineStatusNetworkAuthenticationRequired int32 = 5011 // 语义: 需要网络认证（如wifi认证页面）
	QxEngineStatusParserInitFailed              int32 = 5012 // 语义: 解析器初始化失败
	QxEngineStatusServerEcdhUnSupported         int32 = 5013 // 语义: 服务器不支持ecdh
)
