package hasher

import "hash"

type HashBrute interface {
	Verify(string) bool
	IncrementCount()
	SetKey(string, string)
	GetList() map[string]string
	GetCount() int
	Reset()
	GetHash() hash.Hash
}
