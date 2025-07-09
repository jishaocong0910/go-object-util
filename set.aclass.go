package o

type SetI[T comparable] interface {
	set_() *setM[T]
}

type setM[T comparable] struct {
	i  SetI[T]
	im MapI[T, any]
}

func (this *setM[T]) set_() *setM[T] {
	return this
}

func (this *setM[T]) Add(es ...T) {
	for _, e := range es {
		this.im.map_().Put(e, struct{}{})
	}
}

func (this *setM[T]) AddSet(i SetI[T]) {
	for k := range i.set_().im.map_().m {
		this.im.map_().Put(k, struct{}{})
	}
}

func (this *setM[T]) Remove(e T) bool {
	return this.im.map_().Remove(e)
}

func (this *setM[T]) RemoveAll(e ...T) {
	this.im.map_().RemoveAll(e...)
}

func (this *setM[T]) Contains(ts ...T) bool {
	return this.im.map_().ContainsKeys(ts...)
}

func (this *setM[T]) ContainsAny(ts ...T) bool {
	return this.im.map_().ContainsAnyKey(ts...)
}

func (this *setM[T]) Len() int {
	return this.im.map_().Len()
}

func (this *setM[T]) Empty() bool {
	return this.im.map_().Empty()
}

func (this *setM[T]) Raw() []T {
	return this.im.map_().Keys()
}

func (this *setM[T]) Range(f func(t T)) {
	this.im.map_().Range(func(k T, v any) {
		f(k)
	})
}

func extendSet[T comparable](i SetI[T], im MapI[T, any]) *setM[T] {
	return &setM[T]{i: i, im: im}
}
