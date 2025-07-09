package o

type enumElemI interface {
	enumElem_() *EnumElemM
}

type EnumElemM struct {
	id string
}

func (this *EnumElemM) enumElem_() *EnumElemM {
	return this
}

// ID 枚举值ID，值为枚举集合中的字段名
func (this *EnumElemM) ID() string {
	var id string
	if this != nil {
		id = this.id
	}
	return id
}

// Undefined 是否未定义的枚举
func (this *EnumElemM) Undefined() bool {
	return this == nil
}
