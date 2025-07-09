package o

func NewStrSet(caseSensitive bool, es ...string) *StrSet {
	s := &StrSet{}
	s.setM = extendSet[string](s, NewStrKeyMap[any](caseSensitive))
	s.Add(es...)
	return s
}

type StrSet struct {
	*setM[string]
}
