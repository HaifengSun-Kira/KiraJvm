package reserved

import (
	"gojvm/instructions/base"
	"gojvm/native"
	"gojvm/rtda"
)
import _ "gojvm/native/java/io"
import _ "gojvm/native/java/lang"
import _ "gojvm/native/java/security"
import _ "gojvm/native/java/util/concurrent/atomic"
import _ "gojvm/native/sun/io"
import _ "gojvm/native/sun/misc"
import _ "gojvm/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
