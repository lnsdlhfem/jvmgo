package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type DSUB struct{ base.NoOperandsInstruction }
type FSUB struct{ base.NoOperandsInstruction }
type ISUB struct{ base.NoOperandsInstruction }
type LSUB struct{ base.NoOperandsInstruction }

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	s := v1 - v2
	stack.PushDouble(s)
}
func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	s := v1 - v2
	stack.PushFloat(s)
}
func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := v1 - v2
	stack.PushInt(s)
}
func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := v1 - v2
	stack.PushLong(s)
}
