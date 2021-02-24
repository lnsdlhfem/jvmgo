package loads

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 从局部变量表中获取（double类型）变量，然后推入操作数栈顶
type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

// 索引来自于操作数
func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(self.Index))
}

// 索引来自操作码中
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}
func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}
func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}
func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
