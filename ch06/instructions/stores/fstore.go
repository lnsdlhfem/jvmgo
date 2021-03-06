package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 把变量从操作数栈中弹出，然后存入局部变量表
type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

// 索引来自操作数
func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(self.Index))
}

// 索引来自操作码
func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}
func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}
func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}
func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
