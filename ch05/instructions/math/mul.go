package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DMUL struct{ base.NoOperandsInstruction }
type FMUL struct{ base.NoOperandsInstruction }
type IMUL struct{ base.NoOperandsInstruction }
type LMUL struct{ base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	m := v1 * v2
	stack.PushDouble(m)
}
func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	m := v1 * v2
	stack.PushFloat(m)
}
func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	m := v1 * v2
	stack.PushInt(m)
}
func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	m := v1 * v2
	stack.PushLong(m)
}
