package o

type mapI[K comparable, V any] interface {
	map_() *mapM[K, V]
	key(K) K
}

type mapM[K comparable, V any] struct {
	i mapI[K, V]
	m map[K]*Entry[K, V]
}

func (this *mapM[K, V]) map_() *mapM[K, V] {
	return this
}

func (this *mapM[K, V]) key(k K) K {
	return k
}

func (this *mapM[K, V]) Put(k K, v V) {
	this.m[this.i.key(k)] = &Entry[K, V]{k, v}
}

func (this *mapM[K, V]) PutAll(other mapI[K, V]) {
	for k, v := range other.map_().m {
		this.m[this.i.key(k)] = v
	}
}

func (this *mapM[K, V]) GetEntry(k K) *Entry[K, V] {
	return this.m[this.i.key(k)]
}

func (this *mapM[K, V]) Get(k K) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	}
	return
}

func (this *mapM[K, V]) GetIfAbsent(k K, f func(k K) V) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	} else {
		v = f(k)
		this.Put(k, v)
	}
	return
}

func (this *mapM[K, V]) Remove(k K) bool {
	if this.ContainsKeys(k) {
		delete(this.m, this.i.key(k))
		return true
	}
	return false
}

func (this *mapM[K, V]) RemoveAll(ks ...K) {
	for _, k := range ks {
		delete(this.m, this.i.key(k))
	}
}

func (this *mapM[K, V]) ContainsKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m[this.i.key(k)]; !ok {
			return false
		}
	}
	return true
}

func (this *mapM[K, V]) ContainsAnyKey(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m[this.i.key(k)]; ok {
			return true
		}
	}
	return false
}

func (this *mapM[K, V]) Keys() []K {
	keys := make([]K, 0, len(this.m))
	for _, e := range this.m {
		keys = append(keys, e.Key)
	}
	return keys
}

func (this *mapM[K, V]) Values() []V {
	values := make([]V, 0, len(this.m))
	for _, e := range this.m {
		values = append(values, e.Value)
	}
	return values
}

func (this *mapM[K, V]) Len() int {
	return len(this.m)
}

func (this *mapM[K, V]) Empty() bool {
	return len(this.m) == 0
}

func (this *mapM[K, V]) Raw() map[K]V {
	raw := make(map[K]V, len(this.m))
	for _, e := range this.m {
		raw[e.Key] = e.Value
	}
	return raw
}

func (this *mapM[K, V]) Range(f func(k K, v V)) {
	for _, e := range this.m {
		f(e.Key, e.Value)
	}
}

func extendMap[K comparable, V any](i mapI[K, V]) *mapM[K, V] {
	return &mapM[K, V]{i: i, m: map[K]*Entry[K, V]{}}
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
