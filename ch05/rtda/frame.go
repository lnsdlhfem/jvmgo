package rtda

// 虚拟机栈帧
type Frame struct {
	// 指向下一个栈帧
	lower *Frame
	// 局部变量表
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
}

// 执行方法所需的局部变量表大小和操作数栈深度，是由编译器计算好的，存储在class文件method_info的Code属性中
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
