package rtda

import "jvmgo/ch06/rtda/heap"

// 用于实现局部变量表
type Slot struct {
	// 存放整数
	num int32
	// 存放引用
	ref *heap.Object
}
