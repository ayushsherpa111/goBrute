package hasher

import (
	"hash"
)

type SHA224Brute struct {
	count      int
	hashedList map[string]string
	hashType   hash.Hash
}

func (brt SHA224Brute) Verify(pass string) bool {
	return true
}

func (brt *SHA224Brute) IncrementCount() {
	brt.count += 1
}

func (brt *SHA224Brute) SetKey(key string, value string) {
	(*brt).hashedList[key] = value
}

func (brt *SHA224Brute) GetList() map[string]string {
	return (*brt).hashedList
}

func (brt *SHA224Brute) GetCount() int {
	return (*brt).count
}

func (brt *SHA224Brute) GetHash() hash.Hash {
	return (*brt).hashType
}

func (brt *SHA224Brute) Reset() {
	(*brt).hashType.Reset()
}
