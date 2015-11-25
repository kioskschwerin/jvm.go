package stack

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Pop the top operand stack value
type pop struct{ base.NoOperandsInstruction }

func (self *pop) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type pop2 struct{ base.NoOperandsInstruction }

func (self *pop2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
