package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// 以下指令用于从操作数中获取一个整数（byte、short），扩展为int型，然后推入栈顶
type BIPUSH struct {
	val int8
}
type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
