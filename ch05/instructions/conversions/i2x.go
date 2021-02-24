package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// double类型转换
type I2F struct{ base.NoOperandsInstruction }
type I2D struct{ base.NoOperandsInstruction }
type I2L struct{ base.NoOperandsInstruction }
type I2B struct{ base.NoOperandsInstruction }
type I2C struct{ base.NoOperandsInstruction }
type I2S struct{ base.NoOperandsInstruction }

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}
func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}
func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(uint8(i))
	stack.PushInt(b)
}
func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(uint16(i))
	stack.PushInt(b)
}
func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int16(i))
	stack.PushInt(b)
}
