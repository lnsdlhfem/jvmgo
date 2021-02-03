package classfile

import (
	"encoding/binary"
)

// 类信息读取器
type ClassReader struct {
	data []byte
}

// u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// u2
func (self *ClassReader) readUint16() uint16 {
	// 变量BigEndian，可以从[]byte中解码多字节数据
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	// 返回数组的大小由第一个uint16决定
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
