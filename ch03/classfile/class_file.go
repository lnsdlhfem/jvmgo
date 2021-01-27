package classfile

import "fmt"

// 类文件结构
type ClassFile struct {
	// magic uint32
	// 次版本
	minorVersion uint16
	// 主版本
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 访问标志，标志class文件是类还是接口，是public还是private
	accessFlags uint16
	thisClass   uint16
	superClass  uint16
	interfaces  []uint16
	fields      []*MemberInfo
	methods     []*MemberInfo
	attributes  []AttributeInfo
}

// 将[]byte数据解析为ClassFile对象
func Parse(classData []byte) (cf *ClassFile, err error) {
	// Go没有异常处理机制，只有一个panic-recover机制
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	// 魔数主要用于开头标识文件类型，class文件的魔数是0xCAFEBABE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	// 参考jdk8，只处理45-52的class版本
	switch self.majorVersion {
	case 45:
		return
	// 46之后的版本，次版本号都为0
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// 从常量池中获取类信息
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有java.lang.Object没有超类
}
func (self *ClassFile) Interfaces() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
