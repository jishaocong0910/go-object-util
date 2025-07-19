package o

type Set[T comparable] struct {
	*set__[T]
}

func NewSet[T comparable](es ...T) *Set[T] {
	s := &Set[T]{}
	s.set__ = extendSet[T](s, NewMap[T, any]())
	s.Add(es...)
	return s
}
