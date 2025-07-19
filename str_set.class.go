package o

type StrSet struct {
	*set__[string]
}

func NewStrSet(caseSensitive bool, es ...string) *StrSet {
	s := &StrSet{}
	s.set__ = extendSet[string](s, NewStrKeyMap[any](caseSensitive))
	s.Add(es...)
	return s
}
