package hasher

import (
	"crypto/sha1"
	"encoding/hex"
)

func BeginSHA1Hash(target string) string {
	hasher := sha1.New()
	hasher.Write([]byte(target))
	byteString := hasher.Sum(nil)
	return hex.EncodeToString(byteString)
}
