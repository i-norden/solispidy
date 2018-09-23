package types

type TyVal int

const (
  TY_NIL  TyVal = iota  // This is an empty type, as a default for values with as-of-yet-undetermined types
	TY_INT  TyVal = iota
	TY_STR  TyVal = iota
	TY_BOOL TyVal = iota
	TY_ARR  TyVal = iota
	TY_MAP  TyVal = iota
	TY_FUNC TyVal = iota
	TY_STRC TyVal = iota
)

type TypeNote struct {
	base TyVal
	par0 TyVal
	par1 TyVal
}
