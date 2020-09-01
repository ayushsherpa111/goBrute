package hasher

import (
	"encoding/hex"
	"io"

	"golang.org/x/crypto/md4"
)

func BeginMd4Hash(target string) string {
	var hasher = md4.New()
	io.WriteString(hasher, target)
	return hex.EncodeToString(hasher.Sum(nil))
}
