package bytecode

import (
	"gosh/core/object"
)

type SymbolTable struct {
	parent  *SymbolTable
	store   map[string]int
	symbols []*Symbol
	idx     int
}

func (s *SymbolTable) GetParent() *SymbolTable {
	return s.parent
}
func (s *SymbolTable) GetId() int {
	return s.idx
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store: make(map[string]int),
	}
}

type Symbol struct {
	Name  string
	Index int
	Value *object.Object
}

// Define 返回该符号表idx
func (s *SymbolTable) Define(name string, local bool) int {
	if !local {
		tmp := s
		ok := false
		for tmp != nil {
			_, ok = tmp.store[name]
			if ok {
				return tmp.idx
			}
			tmp = tmp.parent
		}
	}
	_, ok := s.store[name]
	if !ok {
		// 如果不存在相应的符号
		res := &Symbol{
			name,
			len(s.store),
			nil,
		}
		s.store[name] = res.Index
		s.symbols = append(s.symbols, res)
		return s.idx
	}
	return 0
}

func (s *SymbolTable) GetSymbol(idx int) *Symbol {
	return s.symbols[idx]
}

func (s *SymbolTable) GetAllSymbol() []*Symbol {
	return s.symbols
}
