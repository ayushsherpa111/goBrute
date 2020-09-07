package hasher

import "hash"

type HashBrute interface {
	Reset()
	GetHash() hash.Hash
}
