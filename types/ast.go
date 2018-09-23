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
