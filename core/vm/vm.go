package vm

import (
	"github.com/gookit/slog"
	"gosh/compiler/bytecode"
	"gosh/compiler/token"
	"gosh/core/object"
)

type VM struct {
	sp           int
	ip           int
	curSymIdx    int
	constants    []object.Object
	stack        []*object.Object
	SymbolTables []*bytecode.SymbolTable
	instructions []*bytecode.CompilationScope
}

func NewVM(bytecode *bytecode.Bytecode) *VM {
	return &VM{
		sp:           0,
		ip:           0,
		curSymIdx:    0,
		constants:    bytecode.Constants,
		SymbolTables: bytecode.SymbolTables,
		stack:        make([]*object.Object, 2048),
		instructions: bytecode.Instructions,
	}
}

func (v *VM) Run() error {
	slog.Trace("run vm")
	slog.Trace("域总量:", len(v.instructions))
	slog.Trace(v.instructions[0].Instruction)
	for v.ip < len(v.instructions[0].Instruction) {
		switch token.Opcode(v.instructions[0].Instruction[v.ip]) {
		case token.OpConstant:
			// 把常量压进去
			cidx := int(v.instructions[0].Instruction[v.ip+2]) | int(v.instructions[0].Instruction[v.ip+1])<<8
			v.ip += 2
			v.stack[v.sp] = &v.constants[cidx]
			v.sp++
		case token.OpAdd:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpAdd, *right)
			if err != nil {
				return err
			}

			v.stack[v.sp] = &res
			v.sp++
		case token.OpSub:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpSub, *right)
			if err != nil {
				return err
			}

			v.stack[v.sp] = &res
			v.sp++
		case token.OpMul:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpMul, *right)
			if err != nil {
				return err
			}

			v.stack[v.sp] = &res
			v.sp++

		case token.OpDiv:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpDiv, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OpMinus:
			tmp := &object.Int{Value: 0}
			value := v.stack[v.sp-1]
			v.sp--
			res, err := (*tmp).BinaryOp(token.OpSub, *value)
			if err != nil {
				return err
			}

			v.stack[v.sp] = &res
			v.sp++
		}

		v.ip++

	}
	for i := 0; i < v.sp; i++ {
		slog.Info(*v.stack[i])
	}
	return nil
}
