package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func llvmI64(val int64) *constant.Int {
	return constant.NewInt(types.I64, val)
}

func generateLLVM_IR() {
	/*
		set x 5
		set res x
		set name "Jennifer"
	*/
	m := ir.NewModule()
	x := m.NewGlobalDef("x", llvmI64(5))
	_ = m.NewGlobalDef("res", x)
	_ = m.NewGlobalDef("name", constant.NewCharArrayFromString("Jennifer"))

	fmt.Println(m.String())
}
