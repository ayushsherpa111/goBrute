package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

func BeginMd5Hash(target string) string {
	hasher := md5.New()
	hasher.Write([]byte(target))
	byteString := hasher.Sum(nil)
	return hex.EncodeToString(byteString)
}
