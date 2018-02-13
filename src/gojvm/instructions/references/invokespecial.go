package references

import "gojvm/instructions/base"
import "gojvm/rtda"

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}

