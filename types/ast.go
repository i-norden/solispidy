package types

type AST struct {
	Here *interface{}
	Next *interface{}
}

type FnId struct {
	Fnid int64
}

type TyId struct {
	Tyid int64
}

type AssertId struct {
	Assertid int64
}

type CnstInt struct {
	Data [4]uint64
}

type CnstStr struct {
	Data string
}

type CnstBool struct {
	Data bool
}

type FnSymbol struct {
	Symbol string
}

type TySymbol struct {
	Symbol string
}

type SymbolTable struct {
	Fndefs     []AST
	Tydefs     []AST
	Assertdefs []AST
}
