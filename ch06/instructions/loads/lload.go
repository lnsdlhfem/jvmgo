package loads

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 从局部变量表中获取（long类型）变量，然后推入操作数栈顶
type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

// 索引来自于操作数
func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.Index))
}

// 索引来自操作码中
func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}
func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}
func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}
func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
