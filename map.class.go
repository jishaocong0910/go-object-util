package o

func NewMap[K comparable, V any]() *Map[K, V] {
	m := &Map[K, V]{m: map[K]*Entry[K, V]{}}
	m.mapM = extendMap[K, V](m)
	return m
}

type Map[K comparable, V any] struct {
	*mapM[K, V]
	m map[K]*Entry[K, V]
}

func (this *Map[K, V]) key(k K) K {
	return k
}

func (this *Map[K, V]) Put(k K, v V) {
	this.m[this.i.key(k)] = &Entry[K, V]{k, v}
}

func (this *Map[K, V]) PutAll(other MapI[K, V]) {
	other.Range(func(k K, v V) {
		this.Put(k, v)
	})
}

func (this *Map[K, V]) GetEntry(k K) *Entry[K, V] {
	return this.m[this.i.key(k)]
}

func (this *Map[K, V]) Get(k K) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	}
	return
}

func (this *Map[K, V]) GetIfAbsent(k K, f func(k K) V) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	} else {
		v = f(k)
		this.Put(k, v)
	}
	return
}

func (this *Map[K, V]) Remove(k K) bool {
	if this.ContainsKeys(k) {
		delete(this.m, this.i.key(k))
		return true
	}
	return false
}

func (this *Map[K, V]) RemoveAll(ks ...K) {
	for _, k := range ks {
		delete(this.m, this.i.key(k))
	}
}

func (this *Map[K, V]) ContainsKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m[this.i.key(k)]; !ok {
			return false
		}
	}
	return true
}

func (this *Map[K, V]) ContainsAnyKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m[this.i.key(k)]; ok {
			return true
		}
	}
	return false
}

func (this *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(this.m))
	for _, e := range this.m {
		keys = append(keys, e.Key)
	}
	return keys
}

func (this *Map[K, V]) Values() []V {
	values := make([]V, 0, len(this.m))
	for _, e := range this.m {
		values = append(values, e.Value)
	}
	return values
}

func (this *Map[K, V]) Len() int64 {
	return int64(len(this.m))
}

func (this *Map[K, V]) Empty() bool {
	return len(this.m) == 0
}

func (this *Map[K, V]) Raw() map[K]V {
	raw := make(map[K]V, len(this.m))
	for _, e := range this.m {
		raw[e.Key] = e.Value
	}
	return raw
}

func (this *Map[K, V]) Range(f func(k K, v V)) {
	for _, e := range this.m {
		f(e.Key, e.Value)
	}
}
