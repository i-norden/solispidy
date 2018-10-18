package parser

import (
	"errors"
	//"../types"
	"github.com/i-norden/solispidy/types"
)

/*
	The point of the code in this file is to check the AST, and convert it into an
	AST friendlier to type checking and verification.

	Suppose we have the following code:

	(defn foo
		((Uint a) (Uint b)) (Uint)
		(let (x (+ a b))
				 (y (- a b))
			(* x y)))

	As far as the first AST is concerned, this is a bunch of nested functions.
	However, defn is a special function that generates a new function rather than
	doing any real computation. In this case, it creates a function called foo.
	Then the several calls to Uint on the second line specifically are type
	annotations in this context.

	The let function is special as well; all the following expressions with the
	exception of the last one are actually variable bindings. There is no "x" or
	"y" functions in the above code. Instead, the "parameters" in said expressions
	are the expressions that are evaluated, and x and y are the variables they are
	bound to.

	All these special cases (and more) need to be properly covered in order to
	generate correct code.

	In this case, we need to dump all this data into the FnNode type, though some
	changes may need to be made to the exact formatting.

	We can implement this more or less recursively; a contract definition can be
	seen as a correctly formatted header, plus correctly formatted contents, which
	may consist of variables, functions, structs, and assertions, all of which may
	need to be checked for proper formatting. Functions may contain expressions
	internally that need special checking (e.g, let, parameter types, etc.), and
	the same goes for many other things.
*/

// cdr is an old lisp function that returns the "next" value in a list
// cdrlist returns an array of the next val, next of that, etc.
func cdrlist(ast *types.AST) []*types.AST {
	i := ast.Next
	var ret []*types.AST
	for i != nil {
		ret = append(ret, i)
		i = i.Next
	}
	return ret
}

func cdarlist(ast *types.AST) []*types.Symbol {
	i := ast.Next
	var ret []*types.Symbol
	for i != nil {
		ret = append(ret, i.Here)
		i = i.Next
	}
	return ret
}

func nodelist(ast *types.AST) []*types.AST {
	i := ast
	var ret []*types.AST
	for i != nil {
		ret = append(ret, i)
		i = i.Next
	}
	return ret
}

func carlist(ast *types.AST) []*types.Symbol {
	i := ast
	var ret []*types.Symbol
	for i != nil {
		ret = append(ret, i.Here)
		i = i.Next
	}
	return ret
}

func checkGenericNode(ast *types.AST, fnsym string) (*types.AST, error) {
	if ast.Here == nil {
		return nil, errors.New("One of the ASTs is not properly parsed.")
	}
	here := *ast.Here
	if val, ok := here.(types.FnSymbol); ok {
		if val.Symbol == fnsym {
			return ast, nil
		}
	}

	return nil, errors.New("Failed to match pattern")
}

func pullFnSymbol(ast *types.AST) (string, bool) {
	here := *ast.Next.Here
	if val, ok := here.(types.FnSymbol); ok {
		return val.Symbol, ok
	} else {
		return "", false
	}
}

func CheckContractDef(ast *types.AST) (*types.ContractNode, error) {
	if val, err := checkGenericNode(ast, "def-contract"); err != nil {
		here := *val.Here
		here_ := here.(types.Symbol)
		ret := types.ContractNode{(here_.GetLine()), []types.FnNode{}, []types.TyNode{}, []types.VarNode{}, []types.AssertNode{}, []types.FieldNode{}, types.EmptyTable(), ""}
		sym0, _ := pullFnSymbol(ast.Next)
		ret.Symbol = sym0

		// Iterate over defs, typeswitch on defs and add them to contract arrays.
		// If anything defaults, return an error.
		defs := cdrlist(ast)
		for _, def := range defs {
			if node, er := checkFnDef(def); er != nil {
				ret.Funcs = append(ret.Funcs, *node)
			} else if node, er := checkTyDef(def); er != nil {
				ret.Types = append(ret.Types, *node)
			} else {
				return nil, errors.New("Unexpected expression in contract.")
			}
			// Add more cases for asserts and vars.
		}

		// If everything is okay
		return &ret, nil
	}
	return nil, errors.New("Failed to parse contract definition")
}

func checkFnDef(ast *types.AST) (*types.FnNode, error) {
	if val, err := checkGenericNode(ast, "defn"); err != nil {
		here := *val.Here
		here_ := here.(types.Symbol)
		ret := types.FnNode{(here_.GetLine()), "", []types.TypeNote{}, []types.TypeNote{}, true, nil, types.EmptyTable()}
		sym0, _ := pullFnSymbol(ast.Next)
		ret.Symbol = sym0

		// Check if inpars/expars are valid

		// Check if definition is valid

		// Insert parnames into Symbol Table

		// If everything is okay
		return &ret, nil
	}
	return nil, errors.New("Failed to parse function definition")
}

func checkTyDef(ast *types.AST) (*types.TyNode, error) {
	if val, err := checkGenericNode(ast, "defty"); err != nil {
		here := *val.Here
		here_ := here.(types.Symbol)
		ret := types.TyNode{(here_.GetLine()), "", []types.PairNode{}}
		sym0, _ := pullFnSymbol(ast.Next)
		ret.Symbol = sym0

		// Check if fields are valid

		// If everything is okay
		return &ret, nil
	}
	return nil, errors.New("Failed to parse type definition")
}

func checkField(ast *types.AST, tyid string) (*types.FieldNode, error) {
	if ast.Next == nil {
		return nil, errors.New("Expected a field definition with two elements, not one.")
	}

	// This needs to be more complex to handle compound types (mapping, array, etc.).
	here := *ast.Next.Here
	if _, ok := here.(types.TySymbol); ok {
		var ret types.FieldNode
		ret.TyIn = tyid
		ret.TyEx = types.TY_NIL // For now
		if fun, ok := here.(types.FnSymbol); ok {
			ret.Symbol = fun.Symbol
		} else {
			return nil, errors.New("Expected a field definition with a valid field name.")
		}
		if ast.Next.Next != nil {
			return nil, errors.New("Expected a field definition with two elements, not three.")
		}
		return &ret, nil
	}

	return nil, errors.New("Expected a field definition with a valid type.")
}
