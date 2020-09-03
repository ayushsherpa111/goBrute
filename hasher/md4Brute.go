package hasher

import (
	"encoding/hex"
	"io"

	"golang.org/x/crypto/md4"
)

type MD4Brute struct {
	count      int
	hashedList map[string]string
}

func (brt MD4Brute) BeginHash(target string) string {
	var hasher = md4.New()
	io.WriteString(hasher, target)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (brt MD4Brute) Verify(pass string) bool {
	return true
}

func (brt *MD4Brute) IncrementCount() {
	brt.count += 1
}

func (brt *MD4Brute) SetKey(key string, value string) {
	(*brt).hashedList[key] = value
}

func (brt *MD4Brute) GetList() map[string]string {
	return (*brt).hashedList
}

func (brt *MD4Brute) GetCount() int {
	return (*brt).count
}
