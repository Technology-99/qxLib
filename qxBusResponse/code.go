package qxBusResponse

// note: 业务部分状态码 10000-59999

// note: 补充信息性状态码 10000～19999
// note: 补充成功状态码 20000～29999
// note: 补充重定向状态码 30000～39999
// note: 补充客户端错误  40000～49999
// note: 补充服务器错误 50000～59999

const (
	SUCCESS int32 = 10000 + iota
)

const (
	ACCESS_TOKEN_INVALID int32 = iota + 2000
	ACCESS_EXPIRED
	ACCESS_DENY
	ACCESS_NOT_FOUND
	ACCESS_PWD_WRONG
	ACCESS_ACCOUNT_INVALID
	ACCESS_KEY_INVALID
	ACCOUNT_ALREADY_EXISTS
	ACCESS_CODE_WRONG
	GROUP_ALREADY_EXISTS
	ACCESS_TOO_FAST
	DELETE_ADMIN_WRONG
	CANT_CREATE_GROUP
	CANT_CREATE_ACCOUNT
	REFRESH_EXPIRED

	ACCESS_NOT_TRUST_DEVICE

	DATA_EXISTS
	ACCOUNT_NOT_FOUND

	NOT_ALLOW_PARAMS
	SECURITY_TOKEN_INVALID
	SECURITY_TOKEN_EXPIRED
	ACCESS_NOT_WHITE_IP
)

const (
	NOT_FOUND int32 = iota + 3000
)

const (
	FAIL int32 = iota + 4000
	WRONG_PARAM
	NOT_FOUND_METHOD
	METADATA_NOT_FOUND
	AUTHORIZATION_NOT_FOUND
	ACCESSKEY_NOT_FOUND
	WRONG_CAPTCHA
	WECHAT_ERR_USERTOKEN_EXPIRED
	MOVIE_EXIST
	CHAPTER_EXIST
	EPISODE_EXIST
	SCENE_EXIST
	DATA_EXIST
	NOT_INVITED
	ACCESS_DENIED
	NOT_ADMIN
)

const (
	SERVER_WRONG int32 = 5000 + iota
)

const (
	OPERATE_ARTICLE_STATUS_ERR int32 = 6000 + iota
	OPERATE_LABEL_STATUS_ERR
)

const (

	// note: sdk error code %5d
	ERR_INIT_SDK_NOT_CLIENT int32 = 10001 + iota
	ERR_LOGININFO_NIL
	ERR_JSON_MARSHAL
	ERR_INIT_SDK_NOT_LOGINED
)

const (
	// note: dbl 游戏相关
	ERR_SCENE_LOCK int32 = 600001 + iota
	USER_HAS_PAYMENT
	USER_NO_PAYMENT
	USER_PAYMENT_SUCCESS
	USER_PAYMENT_TIMEOUT
	USER_PAYMENT_PROCESSING
	USER_PAYMENT_FAIL
	TRANSACTION_ERR_RESOURCES_OCCUPIED
)

var WrongMessageEn = map[int32]string{
	SUCCESS:                            "success",
	ACCESS_TOKEN_INVALID:               "invalid token",
	ACCESS_EXPIRED:                     "user licence expired",
	REFRESH_EXPIRED:                    "refresh licence expired",
	ACCESS_NOT_TRUST_DEVICE:            "not a trusted device, need secondary verification",
	ACCESS_DENY:                        "permission denied",
	ACCESS_NOT_FOUND:                   "account does not exist",
	ACCESS_PWD_WRONG:                   "incorrect username or password",
	ACCESS_ACCOUNT_INVALID:             "account is disable",
	ACCESS_KEY_INVALID:                 "AccessKey is invalid",
	ACCOUNT_ALREADY_EXISTS:             "user already exists",
	ACCESS_CODE_WRONG:                  "verification code error",
	ACCESS_TOO_FAST:                    "Access too fast",
	GROUP_ALREADY_EXISTS:               "user group already exists",
	DATA_EXISTS:                        "data exists",
	CANT_CREATE_GROUP:                  "Super administrator cannot create groups",
	CANT_CREATE_ACCOUNT:                "unable to create sub-account, please use root account to create one",
	NOT_FOUND:                          "record not found",
	FAIL:                               "fail",
	NOT_FOUND_METHOD:                   "request method not found",
	WRONG_PARAM:                        "param error",
	METADATA_NOT_FOUND:                 "metadata not found",
	AUTHORIZATION_NOT_FOUND:            "authorization not found",
	ACCESSKEY_NOT_FOUND:                "accesskey not found",
	ACCOUNT_NOT_FOUND:                  "account not found",
	WRONG_CAPTCHA:                      "wrong captcha",
	WECHAT_ERR_USERTOKEN_EXPIRED:       "wechat user_token is expired",
	MOVIE_EXIST:                        "movie already exists",
	CHAPTER_EXIST:                      "chapter already exists",
	EPISODE_EXIST:                      "episode already exists",
	SCENE_EXIST:                        "scene already exists",
	DATA_EXIST:                         "data already exists",
	NOT_INVITED:                        "not invited user",
	NOT_ADMIN:                          "not admin user",
	ACCESS_DENIED:                      "access denied",
	DELETE_ADMIN_WRONG:                 "super administrator cannot be deleted",
	SERVER_WRONG:                       "Internal Server Error",
	OPERATE_ARTICLE_STATUS_ERR:         "The article is on the shelf and cannot be operated",
	OPERATE_LABEL_STATUS_ERR:           "Tab is open and not operable",
	ERR_INIT_SDK_NOT_CLIENT:            "sdk client is nil",
	ERR_LOGININFO_NIL:                  "reset time, logininfo is nil",
	ERR_JSON_MARSHAL:                   "json marshal err",
	ERR_INIT_SDK_NOT_LOGINED:           "sdk client isn't logined",
	ERR_SCENE_LOCK:                     "the scene not unlock",
	USER_PAYMENT_SUCCESS:               "payment success",
	USER_PAYMENT_TIMEOUT:               "payment timeout",
	USER_PAYMENT_PROCESSING:            "payment processing",
	USER_PAYMENT_FAIL:                  "payment fail",
	TRANSACTION_ERR_RESOURCES_OCCUPIED: "transaction resources occupied",
	NOT_ALLOW_PARAMS:                   "not allow params",
	SECURITY_TOKEN_INVALID:             "security token invalid",
	SECURITY_TOKEN_EXPIRED:             "security token expired",
	ACCESS_NOT_WHITE_IP:                "access not white ip",
}
