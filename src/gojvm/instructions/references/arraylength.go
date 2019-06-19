package references

import (
	"fmt"
	"gojvm/instructions/base"
	"gojvm/rtda"
)

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	fmt.Printf("Book ARRAY_LEN : %d arrRef : %s\n", arrLen, arrRef.Class().Name())
	stack.PushInt(arrLen)
}
