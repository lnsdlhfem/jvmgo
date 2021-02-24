package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 把变量从操作数栈中弹出，然后存入局部变量表
type DSTORE struct{ base.Index8Instruction }
type DSTORE_0 struct{ base.NoOperandsInstruction }
type DSTORE_1 struct{ base.NoOperandsInstruction }
type DSTORE_2 struct{ base.NoOperandsInstruction }
type DSTORE_3 struct{ base.NoOperandsInstruction }

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

// 索引来自操作数
func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(self.Index))
}

// 索引来自操作码
func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}
func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}
func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}
func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
