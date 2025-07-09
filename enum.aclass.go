package o

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type enumI[E enumElemI] interface {
	enum_() *EnumM[E]
}

type EnumM[E enumElemI] struct {
	elems []E
	idMap map[string]E
}

func (this *EnumM[E]) enum_() *EnumM[E] {
	return this
}

// Elems 返回所有枚举值
func (this *EnumM[E]) Elems() []E {
	var result []E
	if this != nil {
		result = this.elems
	}
	return result
}

// Undefined 返回一个未定义的枚举值
func (this *EnumM[E]) Undefined() E {
	var v E
	return v
}

// OfId 查找ID对应枚举值
func (this *EnumM[E]) OfId(id string) (e E) {
	if this != nil {
		if v, ok := this.idMap[id]; ok {
			e = v
		}
	}
	return
}

// OfIdIgnoreCase 查找ID对应枚举值，不区分大小写
func (this *EnumM[E]) OfIdIgnoreCase(id string) (e E) {
	if this != nil {
		for _, v := range this.elems {
			if strings.EqualFold(v.enumElem_().id, id) {
				return v
			}
		}
	}
	return
}

// Is 判断是否存在指定枚举值
func (this *EnumM[E]) Is(source E, targets ...E) bool {
	if this != nil {
		for _, t := range targets {
			if t.enumElem_().ID() == source.enumElem_().ID() {
				return true
			}
		}
	}
	return false
}

// Not 与Is方法相反
func (this *EnumM[E]) Not(source E, targets ...E) bool {
	return !this.Is(source, targets...)
}

func NewEnum[E enumElemI, ES enumI[E]](e ES) ES {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	if t.Kind() == reflect.Pointer {
		panic("parameter's type must not be a pointer")
	}
	t = reflect.TypeOf(&e).Elem()
	v = reflect.ValueOf(&e).Elem()
	expectedType := reflect.TypeOf((*E)(nil)).Elem()
	v.FieldByName("EnumM").Set(reflect.ValueOf(&EnumM[E]{}))

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		actualType := tf.Type
		if actualType.Kind() == reflect.Pointer {
			actualType = actualType.Elem()
		}
		if !actualType.AssignableTo(expectedType) {
			continue
		}
		if vf.Kind() == reflect.Pointer {
			panic(fmt.Sprintf("%s.%s must not be a pointer type", t.String(), tf.Name))
		}

		var elem E
		evField := vf.FieldByName("EnumElemM")
		if !tf.IsExported() {
			reflect.NewAt(evField.Type(), unsafe.Pointer(evField.UnsafeAddr())).Elem().Set(reflect.ValueOf(&EnumElemM{}))
			elem = reflect.NewAt(vf.Type(), unsafe.Pointer(vf.UnsafeAddr())).Elem().Interface().(E)
		} else {
			evField.Set(reflect.ValueOf(&EnumElemM{}))
			elem = vf.Interface().(E)
		}

		mEv := elem.enumElem_()
		mEv.id = tf.Name

		mE := e.enum_()
		mE.elems = append(mE.elems, elem)
	}

	mE := e.enum_()
	mE.idMap = make(map[string]E, len(mE.elems))
	for _, elem := range mE.elems {
		mE.idMap[elem.enumElem_().id] = elem
	}

	return v.Interface().(ES)
}
