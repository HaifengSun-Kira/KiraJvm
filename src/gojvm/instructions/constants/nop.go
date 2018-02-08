package constants

import "gojvm/instructions/base"
import "gojvm/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {}
