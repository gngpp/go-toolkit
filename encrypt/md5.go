package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(content string) string {
	hash := md5.New()
	hash.Write([]byte(content))
	return hex.EncodeToString(hash.Sum(nil))
}

func MD5Byte(bytes []byte) string {
	return MD5(string(bytes))
}
