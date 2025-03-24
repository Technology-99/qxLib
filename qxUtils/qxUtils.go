package qxUtils

import "github.com/Technology-99/qxLib/qxCrypto"

func GenPassword(realPwd string) (encryptPwdBase, key, salt string) {
	key = qxCrypto.RandStr(32)
	// 生成盐值
	salt = qxCrypto.RandStr(16)
	// 生成hmac-sha256后的密码
	_, basePwd := qxCrypto.HMACSha256(realPwd+salt, key)
	return basePwd, key, salt
}
