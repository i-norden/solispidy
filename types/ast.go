package types

type AST struct {
	Here *interface{}
	Next *AST
}

type SymbolTable struct {
	Fndefs     []AST
	Tydefs     []AST
	Assertdefs []AST
}

// Symbol definitions
type Symbol interface {
	GetLine() int64
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


type FnNode struct{
  Line     int64
  Symbol   int64
  Inpars   []TypeNote
  Expars   []TypeNote
  IsPublic bool
  Def      *Symbol
}

func (f FnNode) GetLine() int64 {
  return f.Line
}

type LetNode struct{
  Line   int64
  Vars   []VarNode
  Def    *Symbol
}

func (l LetNode) GetLine() int64 {
  return l.Line
}

type OpNode struct{
  Line   int64
  Op     Operation
  Pars   []Symbol
}

func (o OpNode) GetLine() int64 {
  return o.Line
}

type PairNode struct{
  A    *Symbol
  B    *Symbol
  Line int64
}

func (p PairNode) GetLine() int64{
  return p.Line
}

type VarNode struct{
	Line   int64
  Symbol int64
  Type   TypeNote
  Def    *Symbol
}

func (v VarNode) GetLine() int64{
  return v.Line
}

type ContractNode struct{
  Line    int64
  Funcs   []FnNode
  Types   []TyNode
  Vars    []VarNode
  Asserts []AssertNode
}

func (c ContractNode) GetLine() int64{
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

type HOFNode struct{
  Line   int64
  Inpars []Symbol
  HoFOp  HigherOrder
}

func (h HOFNode) GetLine() int64{
  return h.Line
}


// ife can be emulated with CondNode: it's just Cond with only one case
type CondNode struct{
  Line    int64
  Cases   []PairNode
  Default Symbol
}

func (c CondNode) GetLine() int64{
  return c.Line
}


type CallNode struct{
  Line   int64
  Symbol int64
  Pars   []Symbol
}

func (c CallNode) GetLine() int64{
  return c.Line
}


type TyNode struct{
	Line   int64
	Symbol int64
	Fields []PairNode
}

func (t TyNode) GetLine() int64{
	return t.Line
}


type AssertNode struct{
	Line int64
	Def  Symbol
}

func (a AssertNode) GetLine() int64{
	return a.Line
}
