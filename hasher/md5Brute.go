package hasher

import (
	"hash"
)

type MD5Brute struct {
	count      int
	hashedList map[string]string
	hashType   hash.Hash
}

func (brt *MD5Brute) Verify(pass string) bool {
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

func (brt *MD5Brute) GetHash() hash.Hash {
	return (*brt).hashType
}

func (brt *MD5Brute) Reset() {
	(*brt).hashType.Reset()
}
