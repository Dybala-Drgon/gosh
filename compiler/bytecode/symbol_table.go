package bytecode

import "gosh/core/object"

type SymbolTable struct {
	parent  *SymbolTable
	store   map[string]int
	symbols []*Symbol
	idx     int
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

func (s *SymbolTable) Define(name string) {
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
	}
}
func (s *SymbolTable) GetSymbol(idx int) *Symbol {
	return s.symbols[idx]
}
func (s *SymbolTable) GetAllSymbol() []*Symbol {
	return s.symbols
}
