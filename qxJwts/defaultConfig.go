package qxJwts

type QxJwtDefaultConfigWithKms struct {
	CertName          string `json:",default=default"`
	CertPublicKeyPath string `json:",default=etc/jwt.public.crt"`
	SignMethod        string `json:",default=ES384"`
	OnlineExp         int    `json:",default=300"`
	DataEncryptName   string `json:",default=default"`
}
