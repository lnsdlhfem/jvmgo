package rtda

// 虚拟机栈帧
type Frame struct {
	// 指向下一个栈帧
	lower *Frame
	// 局部变量表
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
	// 主要用于实现跳转指令
	thread *Thread
	nextPC int
}

// 执行方法所需的局部变量表大小和操作数栈深度，是由编译器计算好的，存储在class文件method_info的Code属性中
func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getter
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}

// setter
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
