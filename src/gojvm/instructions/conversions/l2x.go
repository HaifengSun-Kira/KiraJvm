package conversions

import "gojvm/instructions/base"
import "gojvm/rtda"

type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := int64(l)
	stack.PushLong(d)
}

