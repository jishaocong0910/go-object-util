package o

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	m := &SyncMap[K, V]{}
	m.syncMapM = extendSyncMap[K, V](m)
	return m
}

type SyncMap[K comparable, V any] struct {
	*syncMapM[K, V]
}
