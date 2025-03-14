package qxEngineResponse

// note: qx内部开发引擎部分状态码 1000-5999
// note: 如果是业务上的错误码，请从10000,20000,30000,40000,50000开头

// note: 补充信息性状态码 1000～1999
// note: 补充成功状态码 2000～2999
// note: 补充重定向状态码 3000～3999
// note: 补充客户端错误  4000～4999
// note: 补充服务器错误 5000～5999

const (
	QxEngineStatusContinue                      int32 = 1000 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.2.1
	QxEngineStatusSwitchingProtocols            int32 = 1001 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.2.2
	QxEngineStatusProcessing                    int32 = 1002 // code占用，避免理解歧义, 参考http协议: RFC 2518, 10.1
	QxEngineStatusEarlyHints                    int32 = 1003 // code占用，避免理解歧义, 参考http协议: RFC 8297
	QxEngineStatusAccessTokenInvalid            int32 = 1004 // 错误的访问凭证
	QxEngineStatusAccessExpired                 int32 = 1005 // 访问凭证已过期
	QxEngineStatusRefreshTokenInvalid           int32 = 1006 // 错误的刷新凭证
	QxEngineStatusRefreshTokenExpired           int32 = 1007 // 刷新凭证已过期
	QxEngineStatusAccountNotFound               int32 = 1008 // 找不到该账号
	QxEngineStatusOK                            int32 = 2000 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.1
	QxEngineStatusCreated                       int32 = 2001 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.2
	QxEngineStatusAccepted                      int32 = 2002 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.3
	QxEngineStatusNonAuthoritativeInfo          int32 = 2003 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.4
	QxEngineStatusNoContent                     int32 = 2004 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.5
	QxEngineStatusResetContents                 int32 = 2005 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.6
	QxEngineStatusPartialContent                int32 = 2006 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.3.7
	QxEngineStatusMultiStatus                   int32 = 2007 // code占用，避免理解歧义, 参考http协议: RFC 4918, 11.1
	QxEngineStatusAlreadyReported               int32 = 2008 // code占用，避免理解歧义, 参考http协议: RFC 5842, 7.1
	QxEngineStatusIMUsed                        int32 = 2026 // code占用，避免理解歧义, 参考http协议: RFC 3229, 10.4.1
	QxEngineStatusNotTrustDevice                int32 = 2027 // 不是可信设备
	QxEngineStatusAccountExists                 int32 = 2028 // 账号已经存在
	QxEngineStatusMultipleChoices               int32 = 3000 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.1
	QxEngineStatusMovedPermanently              int32 = 3001 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.2
	QxEngineStatusFound                         int32 = 3002 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.3
	QxEngineStatusSeeOther                      int32 = 3003 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.4
	QxEngineStatusNotModified                   int32 = 3004 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.5
	QxEngineStatusUseProxy                      int32 = 3005 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.6
	QxEngineUnused                              int32 = 3006 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.7 (Unused)
	QxEngineStatusTemporaryRedirect             int32 = 3007 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.8
	QxEngineStatusPermanentRedirect             int32 = 3008 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.4.9
	QxEngineStatusRedirectSuccess               int32 = 3009 // 重定向成功
	QxEngineStatusBadRequest                    int32 = 4000 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.1
	QxEngineStatusUnauthorized                  int32 = 4001 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.2
	QxEngineStatusPaymentRequired               int32 = 4002 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.3
	QxEngineStatusForbidden                     int32 = 4003 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.4
	QxEngineStatusNotFound                      int32 = 4004 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.5
	QxEngineStatusMethodNotAllowed              int32 = 4005 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.6
	QxEngineStatusNotAcceptable                 int32 = 4006 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.7
	QxEngineStatusProxyAuthRequired             int32 = 4007 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.8
	QxEngineStatusRequestTimeout                int32 = 4008 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.9
	QxEngineStatusConflict                      int32 = 4009 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.10
	QxEngineStatusGone                          int32 = 4010 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.11
	QxEngineStatusLengthRequired                int32 = 4011 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.12
	QxEngineStatusPreconditionFailed            int32 = 4012 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.13
	QxEngineStatusRequestEntityTooLarge         int32 = 4013 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.14
	QxEngineStatusRequestURITooLong             int32 = 4014 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.15
	QxEngineStatusUnsupportedMediaType          int32 = 4015 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.16
	QxEngineStatusRequestedRangeNotSatisfiable  int32 = 4016 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.17
	QxEngineStatusExpectationFailed             int32 = 4017 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.18
	QxEngineStatusTeapot                        int32 = 4018 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.19 (Unused)
	QxEngineStatusMisdirectedRequest            int32 = 4021 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.20
	QxEngineStatusUnprocessableEntity           int32 = 4022 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.21
	QxEngineStatusLocked                        int32 = 4023 // code占用，避免理解歧义, 参考http协议: RFC 4918, 11.3
	QxEngineStatusFailedDependency              int32 = 4024 // code占用，避免理解歧义, 参考http协议: RFC 4918, 11.4
	QxEngineStatusTooEarly                      int32 = 4025 // code占用，避免理解歧义, 参考http协议: RFC 8470, 5.2.
	QxEngineStatusUpgradeRequired               int32 = 4026 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.5.22
	QxEngineStatusPreconditionRequired          int32 = 4028 // code占用，避免理解歧义, 参考http协议: RFC 6585, 3
	QxEngineStatusTooManyRequests               int32 = 4029 // code占用，避免理解歧义, 参考http协议: RFC 6585, 4
	QxEngineStatusRequestHeaderFieldsTooLarge   int32 = 4031 // code占用，避免理解歧义, 参考http协议: RFC 6585, 5
	QxEngineStatusUnavailableForLegalReasons    int32 = 4051 // code占用，避免理解歧义, 参考http协议: RFC 7725, 3
	QxEngineStatusAccountOrPasswordInvalid      int32 = 4052 // 账号或者密码错误
	QxEngineStatusInternalServerError           int32 = 5000 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.6.1
	QxEngineStatusNotImplemented                int32 = 5001 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.6.2
	QxEngineStatusBadGateway                    int32 = 5002 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.6.3
	QxEngineStatusServiceUnavailable            int32 = 5003 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.6.4
	QxEngineStatusGatewayTimeout                int32 = 5004 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.6.5
	QxEngineStatusHTTPVersionNotSupported       int32 = 5005 // code占用，避免理解歧义, 参考http协议: RFC 9110, 15.6.6
	QxEngineStatusVariantAlsoNegotiates         int32 = 5006 // code占用，避免理解歧义, 参考http协议: RFC 2295, 8.1
	QxEngineStatusInsufficientStorage           int32 = 5007 // code占用，避免理解歧义, 参考http协议: RFC 4918, 11.5
	QxEngineStatusLoopDetected                  int32 = 5008 // code占用，避免理解歧义, 参考http协议: RFC 5842, 7.2
	QxEngineStatusNotExtended                   int32 = 5010 // code占用，避免理解歧义, 参考http协议: RFC 2774, 7
	QxEngineStatusNetworkAuthenticationRequired int32 = 5011 // code占用，避免理解歧义, 参考http协议: RFC 6585, 6
)
