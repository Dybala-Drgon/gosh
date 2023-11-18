package bytecode

import "gosh/core/object"

type SymbolTable struct {
	parent *SymbolTable
	store  map[string]Symbol
	idx    int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store: make(map[string]Symbol),
	}
}

type Symbol struct {
	Name  string
	Index int
	value *object.Object
}

func (s *SymbolTable) Define(name string) {
	_, ok := s.store[name]
	if !ok {
		// 如果不存在相应的符号
		res := Symbol{
			name,
			len(s.store),
			nil,
		}
		s.store[name] = res
	}
}
