package classfile

// 本身并不存储字符串数据，只存了常量池索引，索引指向一个CONSTANT_Utf8_info常量
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
