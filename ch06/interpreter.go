package main

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/instructions"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 解释器
// 目前只能执行一个java方法
func interpret(methodInfo *classfile.MemberInfo) {
	// 获取方法的Code属性
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	// 创建线程，并创建栈帧推入栈顶
	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)

	// 优雅的结束方法
	defer catchErr(frame)
	// 执行方法
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	// 循环执行“计算pc、解码指令、执行指令”三个步骤，直到遇到错误
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 解码
		reader.Reset(bytecode, pc)
		// 获取操作码
		opcode := reader.ReadUint8()
		// 根据操作码创建对应的指令
		inst := instructions.NewInstruction(opcode)
		// 获取可能的操作数（有的指令不需要操作数）
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// 执行
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
