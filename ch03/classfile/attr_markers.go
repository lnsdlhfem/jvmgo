package classfile

type MarkerAttribute struct{}

// 二者都是标记属性
type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
