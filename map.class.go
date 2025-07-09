package o

func NewMap[K comparable, V any]() *Map[K, V] {
	m := &Map[K, V]{}
	m.mapM = extendMap[K, V](m)
	return m
}

type Map[K comparable, V any] struct {
	*mapM[K, V]
}
