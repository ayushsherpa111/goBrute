package hasher

import (
	"hash"
)

type MD4Brute struct {
	count      int
	hashedList map[string]string
	hashType   hash.Hash
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

func (brt *MD4Brute) GetHash() hash.Hash {
	return (*brt).hashType
}

func (brt *MD4Brute) Reset() {
	(*brt).hashType.Reset()
}
