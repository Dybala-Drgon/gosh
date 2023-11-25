package vm

import (
	"github.com/gookit/slog"
	"gosh/compiler/bytecode"
	"gosh/compiler/token"
	"gosh/core/object"
)

type VM struct {
	sp            int
	ip            int
	curSymIdx     int
	constants     []object.Object
	stack         []*object.Object
	SymbolTables  []*bytecode.SymbolTable
	instructions  []*bytecode.CompilationScope
	functiontable map[string]bytecode.FunInfo
}

func NewVM(bytecode *bytecode.Bytecode) *VM {
	return &VM{
		sp:            0,
		ip:            0,
		curSymIdx:     0,
		constants:     bytecode.Constants,
		SymbolTables:  bytecode.SymbolTables,
		stack:         make([]*object.Object, 2048),
		instructions:  bytecode.Instructions,
		functiontable: bytecode.FuncTable,
	}
}

type Frame struct {
	ip        int
	sp        int
	symIdx    int
	frameType FrameType
}

func (v *VM) Run() error {
	slog.Trace("run vm")
	slog.Trace("域总量:", len(v.instructions))
	var StackFrame []Frame

	//slog.Trace(v.instructions[0].Instruction)
	for v.ip < len(v.instructions[v.curSymIdx].Instruction) {
		switch token.Opcode(v.instructions[v.curSymIdx].Instruction[v.ip]) {
		case token.OpConstant:
			// 把常量压进去
			cidx := int(v.instructions[v.curSymIdx].Instruction[v.ip+2]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+1])<<8
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
		case token.OpLess:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpLess, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OpLessEqual:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpLessEqual, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OPGreater:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OPGreater, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OPGreaterEqual:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OPGreaterEqual, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OpEqual:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpEqual, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OpNotEqual:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpNotEqual, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
			// TODO: finish this
		case token.OpBOr:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpBOr, *right)
			if err != nil {
				return err
			}
			v.stack[v.sp] = &res
			v.sp++
		case token.OpBAnd:
			right := v.stack[v.sp-1]
			left := v.stack[v.sp-2]
			v.sp -= 2

			res, err := (*left).BinaryOp(token.OpBAnd, *right)
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
		case token.OpSetVar:
			symtableidx := int(v.instructions[v.curSymIdx].Instruction[v.ip+2]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+1])<<8
			idx := int(v.instructions[v.curSymIdx].Instruction[v.ip+4]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+3])<<8
			sb := v.SymbolTables[symtableidx].GetSymbol(idx)
			v.ip += 4
			sb.Value = v.stack[v.sp-1]
			v.sp--
		case token.OpGetVar:
			symtableidx := int(v.instructions[v.curSymIdx].Instruction[v.ip+2]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+1])<<8
			idx := int(v.instructions[v.curSymIdx].Instruction[v.ip+4]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+3])<<8
			sb := v.SymbolTables[symtableidx].GetSymbol(idx)
			v.ip += 4
			v.stack[v.sp] = sb.Value
			v.sp++
			//TODO: check this
		case token.OpJumpSymTable:
			symtableidx := int(v.instructions[v.curSymIdx].Instruction[v.ip+2]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+1])<<8
			v.ip += 2
			tmp := Frame{
				ip:        v.ip,
				sp:        v.sp,
				symIdx:    v.curSymIdx,
				frameType: Blk,
			}
			StackFrame = append(StackFrame, tmp)
			v.curSymIdx = symtableidx
			v.ip = -1
		case token.OpExitSymTable:
			length := len(StackFrame)
			tmp := StackFrame[length-1]
			StackFrame = StackFrame[:length-1]
			v.ip = tmp.ip
			v.sp = tmp.sp
			v.curSymIdx = tmp.symIdx
			//v.curSymIdx = v.SymbolTables[v.curSymIdx].GetParent().GetId()
		case token.OpIf:
			stkValue := v.stack[v.sp-1]
			v.sp--
			if (*stkValue).IsFalsy() {
				v.ip += 6
			} else {
				// if  idx  chukou    | else 位置
				// if jsymbol id id jmp chukouid chukouid  jsymbol id id
				// if jsymbol id id jmp chukouid chukouid  if jsymbol id id jmp chukouid chukouid  jsymbol id id
			}
		case token.OpJump:

			exitPos := int(v.instructions[v.curSymIdx].Instruction[v.ip+2]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+1])<<8
			v.ip = exitPos - 1
		case token.OpForIfPre:
			// forpre jmp choukouidx chukouidx jsymbol id id [post stmt] [exp] forifpost jmp pos pos
			//                                  pos
			//
			stkValue := v.stack[v.sp-1]
			v.sp--
			if (*stkValue).IsFalsy() {
				// 如果首轮条件不符合直接跳到指令最后即可
			} else {
				// 进入循环
				v.ip += 3
			}
		case token.OpForIfPost:
			stkValue := v.stack[v.sp-1]
			v.sp--
			if (*stkValue).IsFalsy() {
				v.ip += 3
			} else {

			}
		case token.OpReturn:
			length := len(StackFrame)
			i := length - 1
			for ; i >= 0; i-- {
				tmp := StackFrame[i]
				if tmp.frameType == Func {

					// 需要跳出当前stack frame
					break
				}
			}
			if i == 0 && StackFrame[i].frameType != Func {
				panic("栈帧错误")
			}

			tmp := StackFrame[i]
			StackFrame = StackFrame[:i]
			v.ip = tmp.ip
			v.curSymIdx = tmp.symIdx
		case token.OpCall:
			slog.Info("opcall")

			symtableidx := int(v.instructions[v.curSymIdx].Instruction[v.ip+2]) | int(v.instructions[v.curSymIdx].Instruction[v.ip+1])<<8
			v.ip += 2
			tmp := Frame{
				ip:        v.ip,
				sp:        v.sp,
				symIdx:    v.curSymIdx,
				frameType: Func,
			}
			StackFrame = append(StackFrame, tmp)
			v.curSymIdx = symtableidx
			v.ip = -1
		default:
			slog.Error("尚未支持的符号", token.Opcode(v.instructions[v.curSymIdx].Instruction[v.ip]))
		}
		v.ip++
	}
	for _, elem := range v.SymbolTables[0].GetAllSymbol() {
		slog.Debug(elem.Name, *elem.Value)
	}
	slog.Info("sp = ", v.sp)
	for i := 0; i < v.sp; i++ {
		slog.Info(*v.stack[i])
	}
	slog.Info(v.functiontable)
	return nil
}
