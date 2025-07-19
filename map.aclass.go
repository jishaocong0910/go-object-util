package o

type Map_[K comparable, V any] interface {
	map_() *map__[K, V]
	key(k K) K

	Put(k K, v V)
	PutAll(other Map_[K, V])
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
	NotEmpty() bool
	Raw() map[K]V
	Range(f func(k K, v V))
}

type map__[K comparable, V any] struct {
	i Map_[K, V]
}

func (this *map__[K, V]) map_() *map__[K, V] {
	return this
}

func extendMap[K comparable, V any](i Map_[K, V]) *map__[K, V] {
	return &map__[K, V]{i: i}
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
