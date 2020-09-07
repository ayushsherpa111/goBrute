package hasher

import (
	"hash"
)

type HashStruct struct {
	hashType hash.Hash
}

func (brt *HashStruct) GetHash() hash.Hash {
	return (*brt).hashType
}

func (brt *HashStruct) Reset() {
	(*brt).hashType.Reset()
}
