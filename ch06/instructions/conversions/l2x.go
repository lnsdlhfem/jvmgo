package conversions

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// double类型转换
type L2F struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }
type L2D struct{ base.NoOperandsInstruction }

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}
func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}
