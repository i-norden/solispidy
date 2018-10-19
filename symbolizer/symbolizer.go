package parser

import (
	"errors"
	//"../types"
	symbol "github.com/i-norden/solispidy/common/types"
	ast1 "github.com/i-norden/solispidy/parser/types"
	ast2 "github.com/i-norden/solispidy/symbolizer/types"
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
func cdrlist(ast *ast1.AST) []*ast1.AST {
	i := ast.Next
	var ret []*ast1.AST
	for i != nil {
		ret = append(ret, i)
		i = i.Next
	}
	return ret
}

func cdarlist(ast *ast1.AST) []*symbol.Symbol {
	i := ast.Next
	var ret []*symbol.Symbol
	for i != nil {
		ret = append(ret, i.Here)
		i = i.Next
	}
	return ret
}

func nodelist(ast *ast1.AST) []*ast1.AST {
	i := ast
	var ret []*ast1.AST
	for i != nil {
		ret = append(ret, i)
		i = i.Next
	}
	return ret
}

func carlist(ast *ast1.AST) []*symbol.Symbol {
	i := ast
	var ret []*symbol.Symbol
	for i != nil {
		ret = append(ret, i.Here)
		i = i.Next
	}
	return ret
}

func checkGenericNode(ast *ast1.AST, fnsym string) bool {
	if ast.Here == nil {
		return false
	}
	here := *ast.Here
	if val, ok := here.(ast1.FnSymbol); ok {
		if val.Symbol == fnsym {
			return true
		}
	}

	return false
}

func pullFnSymbol(ast *ast1.AST) (string, bool) {
	here := *ast.Next.Here
	if val, ok := here.(ast1.FnSymbol); ok {
		return val.Symbol, ok
	} else {
		return "", false
	}
}

func CheckFile(asts []ast1.AST) ([]ast2.ContractNode, []error) {
	var retContracts []ast2.ContractNode
	var retErrors []error

	for _, ast := range asts {
		if checkGenericNode(&ast, "def-contract") {
			contract, errs := tryContract(&ast)
			if contract != nil {
				retContracts = append(retContracts, *contract)
			}
			if errs != nil {
				retErrors = append(retErrors, errs...)
			}
		} else {
			retErrors = append(retErrors, errors.New("Improperly defined contract"))
		}
	}

	if len(retErrors) != 0 {
		retContracts = make([]ast2.ContractNode, 0)
	}
	return retContracts, retErrors
}

func tryContract(ast *ast1.AST) (*ast2.ContractNode, []error) {

	var retErrors []error

	// Check contents of contract expression

	return nil, retErrors
}

func checkField(ast *ast1.AST, tyid string) (*ast2.FieldNode, error) {
	if ast.Next == nil {
		return nil, errors.New("Expected a field definition with two elements, not one.")
	}

	// This needs to be more complex to handle compound types (mapping, array, etc.).
	here := *ast.Next.Here
	if _, ok := here.(ast1.TySymbol); ok {
		var ret ast2.FieldNode
		ret.TyIn = tyid
		ret.TyEx = ast2.TY_NIL // For now
		if fun, ok := here.(ast1.FnSymbol); ok {
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