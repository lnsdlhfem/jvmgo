package heap

import "jvmgo/ch06/classfile"

// 将要存放到方法区中的类
type Class struct {
	// 类访问标志
	accessFlag     uint16
	name           string // thisClassName
	superClassName string
	interfaceNames []string
	// 常量池指针
	constantPool *ConstantPool
	// 字段表
	fields []*Field
	// 方法表
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
}

func (self *Class) NewClass(cf *classfile.ClassFile) *Class {
	
}
