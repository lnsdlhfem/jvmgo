package classfile

// 用于存储具体的字符串值，字段名、字段描述符就是以这种格式存储
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	// 调用decodeMUTF8方法将字节数组解析成字符串
	self.str = decodeMUTF8(bytes)
}

// 简化版
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
