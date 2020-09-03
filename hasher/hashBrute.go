package hasher

type HashBrute interface {
	BeginHash(string) string
	Verify(string) bool
	IncrementCount()
	SetKey(string, string)
	GetList() map[string]string
	GetCount() int
}
