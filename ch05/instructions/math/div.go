package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }
type IDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	d := v1 / v2
	stack.PushDouble(d)
}
func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	d := v1 / v2
	stack.PushFloat(d)
}
func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	d := v1 / v2
	stack.PushInt(d)
}
func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	d := v1 / v2
	stack.PushLong(d)
}
