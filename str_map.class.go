package o

import "golang.org/x/text/cases"

func NewStrKeyMap[V any](caseSensitive bool) *StrKeyMap[V] {
	m := &StrKeyMap[V]{caseSensitive: caseSensitive}
	m.mapM = extendMap[string, V](m)
	return m
}

type StrKeyMap[V any] struct {
	*mapM[string, V]
	caseSensitive bool
}

func (this *StrKeyMap[V]) key(k string) string {
	if this.caseSensitive {
		return k
	} else {
		return cases.Fold().String(k)
	}
}
