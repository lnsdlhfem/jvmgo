package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 把变量从操作数栈中弹出，然后存入局部变量表
type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

// 索引来自操作数
func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

// 索引来自操作码
func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}
func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}
func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}
func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
