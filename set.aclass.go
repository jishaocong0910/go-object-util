package o

type SetI[T comparable] interface {
	set_() *setM[T]

	Add(es ...T)
	AddSet(other SetI[T])
	Remove(e T) bool
	RemoveAll(e ...T)
	Contains(es ...T) bool
	ContainsAny(es ...T) bool
	Len() int64
	Empty() bool
	Raw() []T
	Range(f func(e T))
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
		this.im.Put(e, struct{}{})
	}
}

func (this *setM[T]) AddSet(other SetI[T]) {
	other.Range(func(e T) {
		this.Add(e)
	})
}

func (this *setM[T]) Remove(e T) bool {
	return this.im.Remove(e)
}

func (this *setM[T]) RemoveAll(e ...T) {
	this.im.RemoveAll(e...)
}

func (this *setM[T]) Contains(ts ...T) bool {
	return this.im.ContainsKeys(ts...)
}

func (this *setM[T]) ContainsAny(es ...T) bool {
	return this.im.ContainsAnyKeys(es...)
}

func (this *setM[T]) Len() int64 {
	return this.im.Len()
}

func (this *setM[T]) Empty() bool {
	return this.im.Empty()
}

func (this *setM[T]) Raw() []T {
	return this.im.Keys()
}

func (this *setM[T]) Range(f func(e T)) {
	this.im.Range(func(k T, v any) {
		f(k)
	})
}

func extendSet[T comparable](i SetI[T], im MapI[T, any]) *setM[T] {
	return &setM[T]{i: i, im: im}
}
