package o

func NewSyncSet[T comparable](es ...T) *SyncSet[T] {
	s := &SyncSet[T]{}
	s.setM = extendSet[T](s, NewSyncMap[T, any]())
	s.Add(es...)
	return s
}

type SyncSet[T comparable] struct {
	*setM[T]
}
