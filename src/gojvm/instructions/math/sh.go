package math

import "gojvm/instructions/base"
import "gojvm/rtda"

//int left shift
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f // 2^5 == 32
	result := v1 << s
	stack.PushInt(result)
}

//int arithmetic right shift
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f // 2^5 == 32
	result := v1 >> s
	stack.PushInt(result)
}

//int logical right shift
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f // 2^5 == 32
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

//long left shift
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f // 2^6 == 64
	result := v1 << s
	stack.PushLong(result)
}

//long arithmetic right shift
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f // 2^6 == 64
	result := v1 >> s
	stack.PushLong(result)
}

//long logical right shift
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f // 2^6 == 64
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}

