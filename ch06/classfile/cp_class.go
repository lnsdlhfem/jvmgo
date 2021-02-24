package classfile

// 与CONSTANT_String_info类似，本身并不存储字符串数据，只存了常量池索引，索引指向一个CONSTANT_Utf8_info常量
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getUtf8(self.nameIndex)
}
