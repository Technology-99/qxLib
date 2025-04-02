package qxObjectStorage

const (
	ObjectStorageTypeMinio = "minio"
	ObjectStorageTypeCos   = "cos"
	ObjectStorageTypeOss   = "oss"
	ObjectStorageTypeS3    = "s3"
)

var ObjectStorageTypeSupport = map[string]bool{
	ObjectStorageTypeMinio: true,
	ObjectStorageTypeCos:   true,
	ObjectStorageTypeOss:   true,
	ObjectStorageTypeS3:    true,
}

func CheckObjectStorageSupport(osType string) bool {
	if _, ok := ObjectStorageTypeSupport[osType]; ok {
		return true
	} else {
		return false
	}
}
