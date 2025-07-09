package o

type MapI[K comparable, V any] interface {
	map_() *mapM[K, V]
	key(k K) K

	Put(k K, v V)
	PutAll(other MapI[K, V])
	GetEntry(k K) *Entry[K, V]
	Get(k K) (v V)
	GetIfAbsent(k K, f func(k K) V) (v V)
	Remove(k K) bool
	RemoveAll(ks ...K)
	ContainsKeys(ks ...K) bool
	ContainsAnyKeys(ks ...K) bool
	Keys() []K
	Values() []V
	Len() int64
	Empty() bool
	Raw() map[K]V
	Range(f func(k K, v V))
}

type mapM[K comparable, V any] struct {
	i MapI[K, V]
}

func (this *mapM[K, V]) map_() *mapM[K, V] {
	return this
}

func extendMap[K comparable, V any](i MapI[K, V]) *mapM[K, V] {
	return &mapM[K, V]{i: i}
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
