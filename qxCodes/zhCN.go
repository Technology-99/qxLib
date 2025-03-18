package qxCodes

var WrongMessageZhCN = map[int32]string{
	QxEngineStatusContinue:                      "继续。",
	QxEngineStatusSwitchingProtocols:            "服务器同意客户段的协议升级请求",
	QxEngineStatusProcessing:                    "服务器已接收请求，但是仍然在处理中",
	QxEngineStatusEarlyHints:                    "服务器提前发送头部信息（用于优化请求）",
	QxEngineStatusAccessTokenInvalid:            "错误的访问凭证",
	QxEngineStatusAccessExpired:                 "访问凭证已过期",
	QxEngineStatusRefreshTokenInvalid:           "错误的刷新凭证",
	QxEngineStatusRefreshTokenExpired:           "刷新凭证已过期",
	QxEngineStatusAccountNotFound:               "找不到该账号",
	QxEngineStatusDisabled:                      "资源被禁用",
	QxEngineStatusNotSetWhiteList:               "请联系系统管理员设置ip白名单",
	QxEngineStatusNotWhiteList:                  "ip %s, 不在白名单内",
	QxEngineStatusNotFoundAuthorization:         "没找到验证头",
	QxEngineStatusNotFoundXAccessKeyFor:         "没找到api密钥头",
	QxEngineStatusNotFoundMetadata:              "rpc通讯没找到元数据",
	QxEngineStatusOK:                            "请求成功。",
	QxEngineStatusCreated:                       "已创建。",
	QxEngineStatusAccepted:                      "已接受。",
	QxEngineStatusNonAuthoritativeInfo:          "非授权信息。请求成功。但返回的meta信息不在原始的服务器，而是一个副本",
	QxEngineStatusNoContent:                     "无内容。",
	QxEngineStatusResetContents:                 "重置内容。",
	QxEngineStatusPartialContent:                "部分内容。",
	QxEngineStatusMultiStatus:                   "WebDAV扩展，表示多个状态",
	QxEngineStatusAlreadyReported:               "WebDAV扩展，资源已存在",
	QxEngineStatusIMUsed:                        "服务器完成请求，但返回的是变更后的资源",
	QxEngineStatusNotTrustDevice:                "不是可信设备",
	QxEngineStatusAccountExists:                 "账号已经存在",
	QxEngineStatusResourceExists:                "资源已存在",
	QxEngineStatusUnsupportedAlgorithmModel:     "不支持的算法模型",
	QxEngineStatusMultipleChoices:               "多种选择。",
	QxEngineStatusMovedPermanently:              "永久移动。",
	QxEngineStatusFound:                         "临时移动。",
	QxEngineStatusSeeOther:                      "查看其它地址。",
	QxEngineStatusNotModified:                   "未修改。",
	QxEngineStatusUseProxy:                      "使用代理。",
	QxEngineUnused:                              "已废弃",
	QxEngineStatusTemporaryRedirect:             "临时重定向。",
	QxEngineStatusPermanentRedirect:             "永久重定向（类似301，但POST不会变GET）",
	QxEngineStatusRedirectSuccess:               "重定向成功",
	QxEngineStatusBadRequest:                    "客户端请求参数错误",
	QxEngineStatusUnauthorized:                  "请求要求用户的身份认证",
	QxEngineStatusPaymentRequired:               "保留，将来使用",
	QxEngineStatusForbidden:                     "服务器理解请求客户端的请求，但是拒绝执行此请求",
	QxEngineStatusNotFound:                      "服务器无法根据客户端的请求找到资源（网页）。",
	QxEngineStatusMethodNotAllowed:              "客户端请求中的方法被禁止",
	QxEngineStatusNotAcceptable:                 "服务器无法根据客户端请求的内容特性完成请求",
	QxEngineStatusProxyAuthRequired:             "请求要求代理的身份认证，与401类似，但请求者应当使用代理进行授权",
	QxEngineStatusRequestTimeout:                "服务器等待客户端发送的请求时间过长，超时",
	QxEngineStatusConflict:                      "服务器完成客户端的 PUT 请求时可能返回此代码，服务器处理请求时发生了冲突",
	QxEngineStatusGone:                          "客户端请求的资源已经不存在。",
	QxEngineStatusLengthRequired:                "服务器无法处理客户端发送的不带Content-Length的请求信息",
	QxEngineStatusPreconditionFailed:            "客户端请求信息的先决条件错误",
	QxEngineStatusRequestEntityTooLarge:         "由于请求的实体过大，服务器无法处理，因此拒绝请求。为防止客户端的连续请求，服务器可能会关闭连接。如果只是服务器暂时无法处理，则会包含一个Retry-After的响应信息",
	QxEngineStatusRequestURITooLong:             "请求的URI过长（URI通常为网址），服务器无法处理",
	QxEngineStatusUnsupportedMediaType:          "服务器无法处理请求附带的媒体格式",
	QxEngineStatusRequestedRangeNotSatisfiable:  "客户端请求的范围无效",
	QxEngineStatusExpectationFailed:             "服务器无法满足请求头中 Expect 字段指定的预期行为。",
	QxEngineStatusTeapot:                        "愚人节彩蛋",
	QxEngineStatusMisdirectedRequest:            "服务器拒绝处理该请求",
	QxEngineStatusUnprocessableEntity:           "请求格式正确，但数据有误",
	QxEngineStatusLocked:                        "资源被锁定(如WebDAV资源锁)",
	QxEngineStatusFailedDependency:              "由于前一个请求失败，当前请求也无法执行",
	QxEngineStatusTooEarly:                      "服务器拒绝处理可能重放的请求",
	QxEngineStatusUpgradeRequired:               "服务器要求客户端升级协议(如HTTP/1.1->HTTP/2)",
	QxEngineStatusPreconditionRequired:          "需要先满足条件请求（如If-Match头）",
	QxEngineStatusTooManyRequests:               "触发限流",
	QxEngineStatusRequestHeaderFieldsTooLarge:   "请求头字段过大",
	QxEngineStatusUnavailableForLegalReasons:    "由于法律原因，资源不可用",
	QxEngineStatusAccountOrPasswordInvalid:      "账号或者密码错误",
	QxEngineStatusPermissionDenied:              "权限不足",
	QxEngineStatusNotStandardBase64Format:       "不是标准的base64格式",
	QxEngineStatusKmsNotInit:                    "租户的kms尚未完成初始化",
	QxEngineStatusInternalServerError:           "服务器内部错误，无法完成请求",
	QxEngineStatusNotImplemented:                "服务器不支持请求的功能，无法完成请求",
	QxEngineStatusBadGateway:                    "作为网关或者代理工作的服务器尝试执行请求时，从远程服务器接收到了一个无效的响应",
	QxEngineStatusServiceUnavailable:            "由于超载或系统维护，服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的Retry-After头信息中",
	QxEngineStatusGatewayTimeout:                "充当网关或代理的服务器，未及时从远端服务器获取请求",
	QxEngineStatusHTTPVersionNotSupported:       "服务器不支持请求的HTTP协议的版本，无法完成处理",
	QxEngineStatusVariantAlsoNegotiates:         "服务器内部配置错误",
	QxEngineStatusInsufficientStorage:           "服务器存储空间不足",
	QxEngineStatusLoopDetected:                  "服务器监测到无限循环",
	QxEngineStatusNotExtended:                   "服务器需要扩展请求",
	QxEngineStatusNetworkAuthenticationRequired: "需要网络认证（如wifi认证页面）",
	QxEngineStatusParserInitFailed:              "解析器初始化失败",
	QxEngineStatusServerEcdhUnSupported:         "服务器不支持ecdh",
}
