package hasher

import (
	"hash"
)

type HashStruct struct {
	count      int
	hashedList map[string]string
	hashType   hash.Hash
}

func (brt *HashStruct) IncrementCount() {
	brt.count += 1
}

func (brt *HashStruct) SetKey(key string, value string) {
	(*brt).hashedList[key] = value
}

func (brt *HashStruct) GetList() map[string]string {
	return (*brt).hashedList
}

func (brt *HashStruct) GetCount() int {
	return (*brt).count
}

func (brt *HashStruct) GetHash() hash.Hash {
	return (*brt).hashType
}

func (brt *HashStruct) Reset() {
	(*brt).hashType.Reset()
}
