package types

import "github.com/i-norden/solispidy/parser/types"

// Operation definitions

type Operation int64

const (
	OP_ADD Operation = iota
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
	OP_AND
	OP_OR
	OP_XOR
	OP_NOT
	OP_LSS
	OP_GTR
	OP_EQ
	OP_NEQ
	OP_NZR
	OP_ONE
	OP_ZR
)

type FnNode struct {
	Line     int64
	Symbol   string
	Inpars   []TypeNote
	Expars   []TypeNote
	IsPublic bool
	Def      *types.Symbol
	SymTab   SymbolTable
}

type LetNode struct {
	Line   int64
	Vars   []VarNode
	Def    *types.Symbol
	SymTab SymbolTable
}

type OpNode struct {
	Line int64
	Op   Operation
	Pars []types.Symbol
}

type PairNode struct {
	A    *types.Symbol
	B    *types.Symbol
	Line int64
}

type VarNode struct {
	Line   int64
	Symbol string
	Type   TypeNote
	Def    *types.Symbol
}

type ContractNode struct {
	Line    int64
	Funcs   []FnNode
	Types   []TyNode
	Vars    []VarNode
	Asserts []AssertNode
	Fields  []FieldNode
	SymTab  SymbolTable
	Symbol  string
}

type HigherOrder int64

const (
	MAP_HOF   HigherOrder = iota
	REDUC_HOF HigherOrder = iota
	MPRED_HOF HigherOrder = iota
	FILTR_HOF HigherOrder = iota
	LOOP_HOP  HigherOrder = iota
)

type HOFNode struct {
	Line   int64
	Inpars []types.Symbol
	HoFOp  HigherOrder
}

// ife can be emulated with CondNode: it's just Cond with only one case
type CondNode struct {
	Line    int64
	Cases   []PairNode
	Default types.Symbol
}

type CallNode struct {
	Line   int64
	Symbol string
	Pars   []types.Symbol
}

type TyNode struct {
	Line   int64
	Symbol string
	Fields []PairNode
}

type AssertNode struct {
	Line int64
	Def  types.Symbol
}

type FieldNode struct {
	Line   int64
	TyIn   string
	TyEx   TyVal
	Symbol string
}

type SymbolTable struct {
	Types map[string]TypeNote
	Count int64
}

func EmptyTable() SymbolTable {
	var ret SymbolTable
	ret.Types = map[string]TypeNote{}
	ret.Count = 0
	return ret
}

type Scope struct {
	Stack []*SymbolTable
}

type DefTable struct {
	FnDefs map[int64]FnNode
	TyDefs map[int64]TyNode
	Fields map[int64]([]FieldNode)
	VrDefs map[int64]VarNode
	AsDefs map[int64]AssertNode
}

func (f *FnNode) GetLine() int64 {
	return f.Line
}

func (l *LetNode) GetLine() int64 {
	return l.Line
}

func (o *OpNode) GetLine() int64 {
	return o.Line
}

func (p *PairNode) GetLine() int64 {
	return p.Line
}

func (v *VarNode) GetLine() int64 {
	return v.Line
}

func (c *ContractNode) GetLine() int64 {
	return c.Line
}

func (h *HOFNode) GetLine() int64 {
	return h.Line
}

func (c *CondNode) GetLine() int64 {
	return c.Line
}

func (c *CallNode) GetLine() int64 {
	return c.Line
}

func (t *TyNode) GetLine() int64 {
	return t.Line
}

func (a *AssertNode) GetLine() int64 {
	return a.Line
}

func (f *FieldNode) GetLine() int64 {
	return f.Line
}
