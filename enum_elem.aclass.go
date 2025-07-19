package o

type enumElem_ interface {
	enumElem_() *EnumElem__
}

type EnumElem__ struct {
	id string
}

func (this *EnumElem__) enumElem_() *EnumElem__ {
	return this
}

// ID 枚举值ID，值为枚举集合中的字段名
func (this *EnumElem__) ID() string {
	var id string
	if this != nil {
		id = this.id
	}
	return id
}

// Undefined 是否未定义的枚举
func (this *EnumElem__) Undefined() bool {
	return this == nil
}
