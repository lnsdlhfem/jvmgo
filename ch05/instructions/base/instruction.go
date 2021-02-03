package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 跳转指令
type BranchInstruction struct {
	// 跳转偏移量
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 局部变量表索引类指令
type Index8Instruction struct {
	// 局部变量表索引
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 运行时常量池索引类指令
type Index16Instruction struct {
	// 运行时常量池索引
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
