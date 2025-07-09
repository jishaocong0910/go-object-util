package o

func NewSet[T comparable](es ...T) *Set[T] {
	s := &Set[T]{}
	s.setM = extendSet[T](s, NewMap[T, any]())
	s.Add(es...)
	return s
}

type Set[T comparable] struct {
	*setM[T]
}
