package hasher

import "hash"

type SHA1Brute struct {
	count      int
	hashedList map[string]string
	hashType   hash.Hash
}

func (brt *SHA1Brute) Verify(pass string) bool {
	return true
}

func (brt *SHA1Brute) IncrementCount() {
	brt.count += 1
}

func (brt *SHA1Brute) SetKey(key string, value string) {
	(*brt).hashedList[key] = value
}

func (brt *SHA1Brute) GetList() map[string]string {
	return (*brt).hashedList
}

func (brt *SHA1Brute) GetCount() int {
	return (*brt).count
}

func (brt *SHA1Brute) GetHash() hash.Hash {
	return (*brt).hashType
}

func (brt *SHA1Brute) Reset() {
	(*brt).hashType.Reset()
}
