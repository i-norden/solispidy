package types

// Symbol interface definition
type Symbol interface {
	GetLine() int64
}

// Basic AST structure
type AST struct {
	Here Symbol
	Next *AST
}

// Basic AST symbols
type LeftPar struct {
	LPId int64
	Line int64
}

type RightPar struct {
	RPId int64
	Line int64
}

type FnId struct {
	Fnid int64
	Line int64
}

type TyId struct {
	Tyid int64
	Line int64
}

type AssertId struct {
	Assertid int64
	Line     int64
}

type CnstInt struct {
	Data [4]uint64
	Line int64
}

type CnstStr struct {
	Data string
	Line int64
}

type CnstBool struct {
	Data bool
	Line int64
}

type FnSymbol struct {
	Symbol string
	Line   int64
}

type TySymbol struct {
	Symbol string
	Line   int64
}

// GetLine method definitions

func (a AST) GetLine() int64 {
	if a.Here == nil {
		return -1
	}

	return a.Here.GetLine()
}

func (f *LeftPar) GetLine() int64 {
	return f.Line
}

func (f *RightPar) GetLine() int64 {
	return f.Line
}

func (f *FnId) GetLine() int64 {
	return f.Line
}

func (f *TyId) GetLine() int64 {
	return f.Line
}

func (f *AssertId) GetLine() int64 {
	return f.Line
}

func (f *CnstInt) GetLine() int64 {
	return f.Line
}

func (f *CnstStr) GetLine() int64 {
	return f.Line
}

func (f *CnstBool) GetLine() int64 {
	return f.Line
}

func (f *FnSymbol) GetLine() int64 {
	return f.Line
}

func (f *TySymbol) GetLine() int64 {
	return f.Line
}
