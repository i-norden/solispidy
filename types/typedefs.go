package types

type TyVal int64

const (
	TY_NIL TyVal = iota // This is an empty type, as a default for values with as-of-yet-undetermined types
	TY_INT
	TY_STR
	TY_BOOL
	TY_ARR
	TY_MAP
	TY_FUNC
	TY_STRC
)

type TypeNote struct {
	base TyVal
	par0 TyVal
	par1 TyVal
}
