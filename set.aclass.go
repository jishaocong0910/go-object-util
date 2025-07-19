package o

type Set_[T comparable] interface {
	set_() *set__[T]

	Add(es ...T)
	AddSet(other Set_[T])
	Remove(e T) bool
	RemoveAll(e ...T)
	Contains(es ...T) bool
	ContainsAny(es ...T) bool
	Len() int64
	Empty() bool
	NotEmpty() bool
	Raw() []T
	Range(f func(e T))
}

type set__[T comparable] struct {
	i  Set_[T]
	im Map_[T, any]
}

func (this *set__[T]) set_() *set__[T] {
	return this
}

func (this *set__[T]) Add(es ...T) {
	for _, e := range es {
		this.im.Put(e, struct{}{})
	}
}

func (this *set__[T]) AddSet(other Set_[T]) {
	other.Range(func(e T) {
		this.Add(e)
	})
}

func (this *set__[T]) Remove(e T) bool {
	return this.im.Remove(e)
}

func (this *set__[T]) RemoveAll(e ...T) {
	this.im.RemoveAll(e...)
}

func (this *set__[T]) Contains(ts ...T) bool {
	return this.im.ContainsKeys(ts...)
}

func (this *set__[T]) ContainsAny(es ...T) bool {
	return this.im.ContainsAnyKeys(es...)
}

func (this *set__[T]) Len() int64 {
	return this.im.Len()
}

func (this *set__[T]) Empty() bool {
	return this.im.Empty()
}

func (this *set__[T]) NotEmpty() bool {
	return this.im.NotEmpty()
}

func (this *set__[T]) Raw() []T {
	return this.im.Keys()
}

func (this *set__[T]) Range(f func(e T)) {
	this.im.Range(func(k T, v any) {
		f(k)
	})
}

func extendSet[T comparable](i Set_[T], im Map_[T, any]) *set__[T] {
	return &set__[T]{i: i, im: im}
}
