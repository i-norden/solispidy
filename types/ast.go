package types

type AST struct {
	Here *Symbol
	Next *AST
}

// Symbol definitions
type Symbol interface {
	GetLine() int64
}

type LeftPar struct {
	LPId int64
	Line int64
}

func (f LeftPar) GetLine() int64 {
	return f.Line
}

type RightPar struct {
	RPId int64
	Line int64
}

func (f RightPar) GetLine() int64 {
	return f.Line
}

type FnId struct {
	Fnid int64
	Line int64
}

func (f FnId) GetLine() int64 {
	return f.Line
}

type TyId struct {
	Tyid int64
	Line int64
}

func (f TyId) GetLine() int64 {
	return f.Line
}

type AssertId struct {
	Assertid int64
	Line     int64
}

func (f AssertId) GetLine() int64 {
	return f.Line
}

type CnstInt struct {
	Data [4]uint64
	Line int64
}

func (f CnstInt) GetLine() int64 {
	return f.Line
}

type CnstStr struct {
	Data string
	Line int64
}

func (f CnstStr) GetLine() int64 {
	return f.Line
}

type CnstBool struct {
	Data bool
	Line int64
}

func (f CnstBool) GetLine() int64 {
	return f.Line
}

type FnSymbol struct {
	Symbol string
	Line   int64
}

func (f FnSymbol) GetLine() int64 {
	return f.Line
}

type TySymbol struct {
	Symbol string
	Line   int64
}

func (f TySymbol) GetLine() int64 {
	return f.Line
}

type Operation int64

const (
	OP_ADD Operation = iota
	OP_SUB Operation = iota
	OP_MUL Operation = iota
	OP_DIV Operation = iota
	OP_MOD Operation = iota
	OP_AND Operation = iota
	OP_OR  Operation = iota
	OP_XOR Operation = iota
	OP_NOT Operation = iota
	OP_LSS Operation = iota
	OP_GTR Operation = iota
	OP_EQ  Operation = iota
	OP_NEQ Operation = iota
	OP_NZR Operation = iota
	OP_ONE Operation = iota
	OP_ZR  Operation = iota
)

type FnNode struct {
	Line     int64
	Symbol   string
	Inpars   []TypeNote
	Expars   []TypeNote
	IsPublic bool
	Def      *Symbol
	SymTab   SymbolTable
}

func (f FnNode) GetLine() int64 {
	return f.Line
}

type LetNode struct {
	Line   int64
	Vars   []VarNode
	Def    *Symbol
	SymTab SymbolTable
}

func (l LetNode) GetLine() int64 {
	return l.Line
}

type OpNode struct {
	Line int64
	Op   Operation
	Pars []Symbol
}

func (o OpNode) GetLine() int64 {
	return o.Line
}

type PairNode struct {
	A    *Symbol
	B    *Symbol
	Line int64
}

func (p PairNode) GetLine() int64 {
	return p.Line
}

type VarNode struct {
	Line   int64
	Symbol string
	Type   TypeNote
	Def    *Symbol
}

func (v VarNode) GetLine() int64 {
	return v.Line
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

func (c ContractNode) GetLine() int64 {
	return c.Line
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
	Inpars []Symbol
	HoFOp  HigherOrder
}

func (h HOFNode) GetLine() int64 {
	return h.Line
}

// ife can be emulated with CondNode: it's just Cond with only one case
type CondNode struct {
	Line    int64
	Cases   []PairNode
	Default Symbol
}

func (c CondNode) GetLine() int64 {
	return c.Line
}

type CallNode struct {
	Line   int64
	Symbol string
	Pars   []Symbol
}

func (c CallNode) GetLine() int64 {
	return c.Line
}

type TyNode struct {
	Line   int64
	Symbol string
	Fields []PairNode
}

func (t TyNode) GetLine() int64 {
	return t.Line
}

type AssertNode struct {
	Line int64
	Def  Symbol
}

func (a AssertNode) GetLine() int64 {
	return a.Line
}

type FieldNode struct {
	Line   int64
	TyIn   int64
	TyEx   int64
	Symbol string
}

func (f FieldNode) GetLine() int64 {
	return f.Line
}

type SymbolTable struct {
	SymsToIds map[string]int64
	IdsToSyms map[int64]string
	Types     map[int64]TypeNote
	Count     int64
}

func EmptyTable() SymbolTable {
	return SymbolTable{map[string]int64{}, map[int64]string{}, map[int64]TypeNote{}, 0}
}

func (s SymbolTable) addSym(sym string) int64 {
	if val, ok := s.SymsToIds[sym]; ok {
		return val
	} else {
		s.Count++
		s.SymsToIds[sym] = s.Count
		s.IdsToSyms[s.Count] = sym
		s.Types[s.Count] = TypeNote{TY_NIL, TY_NIL, TY_NIL}
		return s.Count
	}
}

func (s SymbolTable) getSym(id int64) (string, bool) {
	if val, ok := s.IdsToSyms[id]; ok {
		return val, ok
	} else {
		return "", false
	}
}

func (s SymbolTable) getId(sym string) (int64, bool) {
	if val, ok := s.SymsToIds[sym]; ok {
		return val, ok
	} else {
		return -1, false
	}
}

func (s SymbolTable) getType(id int64) (TypeNote, bool) {
	if val, ok := s.Types[id]; ok {
		return val, ok
	} else {
		return TypeNote{TY_NIL, TY_NIL, TY_NIL}, false
	}
}

func (s SymbolTable) setType(id int64, t TypeNote) bool {
	if _, ok := s.Types[id]; ok {
		s.Types[id] = t
		return true
	}
	return false
}

type Scope struct {
	Stack []*SymbolTable
}

func (s Scope) searchId(sym string) (int64, bool) {
	for i := len(s.Stack) - 1; i >= 0; i++ {
		if val, ok := s.Stack[i].getId(sym); ok {
			return val, ok
		}
	}
	return -1, false
}

func (s Scope) searchSym(id int64) (string, bool) {
	for i := len(s.Stack) - 1; i >= 0; i++ {
		if val, ok := s.Stack[i].getSym(id); ok {
			return val, ok
		}
	}
	return "", false
}

func (s Scope) searchType(id int64) (TypeNote, bool) {
	for i := len(s.Stack) - 1; i >= 0; i++ {
		if val, ok := s.Stack[i].getType(id); ok {
			return val, ok
		}
	}
	return TypeNote{TY_NIL, TY_NIL, TY_NIL}, false
}

func (s Scope) setType(id int64, t TypeNote) bool {
	for i := len(s.Stack) - 1; i >= 0; i++ {
		if ok := s.Stack[i].setType(id, t); ok {
			return ok
		}
	}
	return false
}

func (s Scope) addTable(syms *SymbolTable) {
	s.Stack = append(s.Stack, syms)
}

type DefTable struct {
	FnDefs map[int64]FnNode
	TyDefs map[int64]TyNode
	Fields map[int64]([]FieldNode)
	VrDefs map[int64]VarNode
	AsDefs map[int64]AssertNode
}
