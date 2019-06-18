package references

import (
	"gojvm/instructions/base"
	"gojvm/rtda"
	"gojvm/rtda/heap"
)

type ANEW_ARRAY struct{
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef :=  cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
