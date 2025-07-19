package o

type SyncSet[T comparable] struct {
	*set__[T]
}

func NewSyncSet[T comparable](es ...T) *SyncSet[T] {
	s := &SyncSet[T]{}
	s.set__ = extendSet[T](s, NewSyncMap[T, any]())
	s.Add(es...)
	return s
}
