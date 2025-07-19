package o

import "golang.org/x/text/cases"

type StrKeyMap[V any] struct {
	*Map[string, V]
	caseSensitive bool
}

func (this *StrKeyMap[V]) key(k string) string {
	if this.caseSensitive {
		return k
	} else {
		return cases.Fold().String(k)
	}
}

func NewStrKeyMap[V any](caseSensitive bool) *StrKeyMap[V] {
	m := &StrKeyMap[V]{Map: NewMap[string, V](), caseSensitive: caseSensitive}
	m.map__ = extendMap[string, V](m)
	return m
}
