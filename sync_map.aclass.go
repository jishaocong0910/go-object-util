package o

import (
	"sync"
	"sync/atomic"
)

type syncMapI[K comparable, V any] interface {
	syncMapM_() *syncMapM[K, V]
	key(K) K
}

type syncMapM[K comparable, V any] struct {
	i   syncMapI[K, V]
	m   sync.Map
	len int64
}

func (this *syncMapM[K, V]) syncMapM_() *syncMapM[K, V] {
	return this
}

func (this *syncMapM[K, V]) key(k K) K {
	return k
}

func (this *syncMapM[K, V]) Put(k K, v V) {
	this.m.Store(this.i.key(k), &Entry[K, V]{k, v})
	atomic.AddInt64(&this.len, 1)
}

func (this *syncMapM[K, V]) PutAll(other syncMapI[K, V]) {
	other.syncMapM_().Range(func(k K, v V) {
		this.Put(k, v)
	})
}

func (this *syncMapM[K, V]) GetEntry(k K) *Entry[K, V] {
	if value, ok := this.m.Load(this.i.key(k)); ok {
		return value.(*Entry[K, V])
	}
	return nil
}

func (this *syncMapM[K, V]) Get(k K) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	}
	return
}

func (this *syncMapM[K, V]) GetIfAbsent(k K, f func(k K) V) (v V) {
	value, _ := this.m.LoadOrStore(this.i.key(k), f(k))
	return value
}

func (this *syncMapM[K, V]) Remove(k K) bool {
	_, exists := this.m.LoadAndDelete(this.i.key(k))
	atomic.AddInt64(&this.len, -1)
	return exists
}

func (this *syncMapM[K, V]) RemoveAll(ks ...K) {
	for _, k := range ks {
		this.Remove(k)
	}
}

func (this *syncMapM[K, V]) ContainsKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m.Load(this.i.key(k)); !ok {
			return false
		}
	}
	return true
}

func (this *syncMapM[K, V]) ContainsAnyKey(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m.Load(this.i.key(k)); ok {
			return true
		}
	}
	return false
}

func (this *syncMapM[K, V]) Keys() []K {
	keys := make([]K, this.len)
	this.m.Range(func(_, v any) bool {
		e := v.(*Entry[K, V])
		keys = append(keys, e.Key)
		return true
	})
	return keys
}

func (this *syncMapM[K, V]) Values() []V {
	values := make([]V, this.len)
	this.m.Range(func(_, v any) bool {
		e := v.(*Entry[K, V])
		values = append(values, e.Value)
		return true
	})
	return values
}

func (this *syncMapM[K, V]) Len() int64 {
	return this.len
}

func (this *syncMapM[K, V]) Empty() bool {
	return this.len == 0
}

func (this *syncMapM[K, V]) Raw() map[K]V {
	raw := make(map[K]V, this.len)
	this.m.Range(func(k, v any) bool {
		e := v.(*Entry[K, V])
		raw[e.Key] = e.Value
		return true
	})
	return raw
}

func (this *syncMapM[K, V]) Range(f func(k K, v V)) {
	this.m.Range(func(_, v any) bool {
		e := v.(*Entry[K, V])
		f(e.Key, e.Value)
		return true
	})
}

func extendSyncMap[K comparable, V any](i syncMapI[K, V]) *syncMapM[K, V] {
	return &syncMapM[K, V]{i: i}
}
