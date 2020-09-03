package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5Brute struct {
	count      int
	hashedList map[string]string
}

func (brt MD5Brute) BeginHash(target string) string {
	var hasher = md5.New()
	hasher.Write([]byte(target))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (brt MD5Brute) Verify(pass string) bool {
	return true
}

func (brt *MD5Brute) IncrementCount() {
	brt.count += 1
}

func (brt *MD5Brute) SetKey(key string, value string) {
	(*brt).hashedList[key] = value
}

func (brt *MD5Brute) GetList() map[string]string {
	return (*brt).hashedList
}

func (brt *MD5Brute) GetCount() int {
	return (*brt).count
}
