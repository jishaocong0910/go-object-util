package o

import (
	"sync"
	"sync/atomic"
)

type SyncMap[K comparable, V any] struct {
	*mapM[K, V]
	m   sync.Map
	len int64
}

func (this *SyncMap[K, V]) syncMapM_() *SyncMap[K, V] {
	return this
}

func (this *SyncMap[K, V]) key(k K) K {
	return k
}

func (this *SyncMap[K, V]) Put(k K, v V) {
	this.m.Store(this.key(k), &Entry[K, V]{k, v})
	atomic.AddInt64(&this.len, 1)
}

func (this *SyncMap[K, V]) PutAll(other MapI[K, V]) {
	other.Range(func(k K, v V) {
		this.Put(k, v)
	})
}

func (this *SyncMap[K, V]) GetEntry(k K) *Entry[K, V] {
	if value, ok := this.m.Load(this.key(k)); ok {
		return value.(*Entry[K, V])
	}
	return nil
}

func (this *SyncMap[K, V]) Get(k K) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	}
	return
}

func (this *SyncMap[K, V]) GetIfAbsent(k K, f func(k K) V) (v V) {
	value, exists := this.m.LoadOrStore(this.key(k), f(k))
	if !exists {
		atomic.AddInt64(&this.len, 1)
	}
	return value.(V)
}

func (this *SyncMap[K, V]) Remove(k K) bool {
	_, exists := this.m.LoadAndDelete(this.key(k))
	if exists {
		atomic.AddInt64(&this.len, -1)
	}
	return exists
}

func (this *SyncMap[K, V]) RemoveAll(ks ...K) {
	for _, k := range ks {
		this.Remove(k)
	}
}

func (this *SyncMap[K, V]) ContainsKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m.Load(this.key(k)); !ok {
			return false
		}
	}
	return true
}

func (this *SyncMap[K, V]) ContainsAnyKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m.Load(this.key(k)); ok {
			return true
		}
	}
	return false
}

func (this *SyncMap[K, V]) Keys() []K {
	keys := make([]K, this.len)
	this.m.Range(func(_, v any) bool {
		e := v.(*Entry[K, V])
		keys = append(keys, e.Key)
		return true
	})
	return keys
}

func (this *SyncMap[K, V]) Values() []V {
	values := make([]V, this.len)
	this.m.Range(func(_, v any) bool {
		e := v.(*Entry[K, V])
		values = append(values, e.Value)
		return true
	})
	return values
}

func (this *SyncMap[K, V]) Len() int64 {
	return this.len
}

func (this *SyncMap[K, V]) Empty() bool {
	return this.len == 0
}

func (this *SyncMap[K, V]) Raw() map[K]V {
	raw := make(map[K]V, this.len)
	this.m.Range(func(k, v any) bool {
		e := v.(*Entry[K, V])
		raw[e.Key] = e.Value
		return true
	})
	return raw
}

func (this *SyncMap[K, V]) Range(f func(k K, v V)) {
	this.m.Range(func(_, v any) bool {
		e := v.(*Entry[K, V])
		f(e.Key, e.Value)
		return true
	})
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	m := &SyncMap[K, V]{}
	m.mapM = extendMap[K, V](m)
	return m
}
