package parser

import (
	"errors"
	//"../types"
	ast1 "github.com/i-norden/solispidy/parser/types"
	ast2 "github.com/i-norden/solispidy/symbolizer/types"
	"github.com/i-norden/solispidy/common/utils"
	"fmt"
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


func checkGenericNode(ast *ast1.AST, fnsym string) bool {
	if ast.Here == nil {
		return false
	}
	//here := ast.Here
	if val, ok := ast.Here.(*ast1.FnSymbol); ok {
		if val.Symbol == fnsym {
			return true
		}
	}else if val, ok := ast.Here.(*ast1.TySymbol); ok {
		if val.Symbol == fnsym {
			return true
		}
	}

	return false
}



func pullFnSymbol(ast *ast1.AST) (string, bool) {
	if val, ok := ast.Next.Here.(*ast1.FnSymbol); ok {
		return val.Symbol, ok
	} else {
		return "", false
	}
}



func CheckFile(asts map[string]ast1.AST) ([]ast2.ContractNode, []error) {
	var retContracts []ast2.ContractNode
	var retErrors []error

	for fname, ast := range asts {
		if checkGenericNode(&ast, "def-contract") {
			contract, errs := tryContract(&ast)
			if contract != nil {
				retContracts = append(retContracts, *contract)
			}
			if errs != nil {
				retErrors = append(retErrors, fmt.Errorf("\nIn file: %s", fname))
				retErrors = append(retErrors, errs...)
			}
		} else {
			retErrors = append(retErrors, fmt.Errorf("\nIn file: %s", fname))
			retErrors = append(retErrors, utils.LineError(ast.GetLine(), "Improperly defined contract"))
		}
	}

	if len(retErrors) != 0 {
		retContracts = make([]ast2.ContractNode, 0)
	}
	return retContracts, retErrors
}



func tryContract(ast *ast1.AST) (*ast2.ContractNode, []error) {

	var retErrors []error
	var retContract ast2.ContractNode

	// Check contents of contract expression
	if ast.Next != nil {
		if val, ok := ast.Next.Here.(*ast1.TySymbol); ok {
			retContract.Symbol = val.Symbol
			retContract.Line   = ast.GetLine()

			// Check internal definitions
			def := ast
			//nilast := ast1.AST{Next: nil, Here: nil}
			for def.Next != nil && def.Next.Here != nil {
				def = def.Next
				if vl, ok := def.Here.(*ast1.AST); !ok || def.Here == nil {
					retErrors = append(retErrors, utils.LineError(def.GetLine(), "Expected a definition (list) here."))
				}else if fn, er := tryFunc(vl); er == nil{
					retContract.Funcs = append(retContract.Funcs, *fn)
				}else if ty, er := tryType(vl); er == nil{
					retContract.Types = append(retContract.Types, *ty)
				}else if vr, er := tryPublic(vl); er == nil{
					retContract.Vars = append(retContract.Vars, *vr)
				}else{
					retErrors = append(retErrors, utils.LineError(def.GetLine(), "Expression not a function, type, assertion, or variable definition."))
				}

			}

			return &retContract, retErrors
		}else{
			retErrors = append(retErrors, errors.New("Contract has no valid name."))
		}
	}else{
		retErrors = append(retErrors, errors.New("Contract header has no contents."))
	}

	return nil, retErrors
}



func tryFunc(ast *ast1.AST) (*ast2.FnNode, []error) {

	var retErrors []error
	var retFunc   ast2.FnNode

	if !checkGenericNode(ast, "defn") {
		retErrors = append(retErrors, utils.LineError(ast.GetLine(), "Improperly defined function"))
		return nil, retErrors
	}

	// Check contents of contract expression
	if ast.Next != nil {
		if val, ok := ast.Next.Here.(*ast1.FnSymbol); ok {
			retFunc.Symbol = val.Symbol
			retFunc.Line   = ast.GetLine()

			// Check internal definitions

			return &retFunc, retErrors
		}else{
			retErrors = append(retErrors, errors.New("Function has no valid name."))
		}
	}else{
		retErrors = append(retErrors, errors.New("Function header has no contents."))
	}

	return nil, retErrors
}



func tryType(ast *ast1.AST) (*ast2.TyNode, []error) {

	var retErrors []error
	var retType   ast2.TyNode

	if !checkGenericNode(ast, "defty") {
		retErrors = append(retErrors, utils.LineError(ast.GetLine(), "Improperly defined struct"))
		return nil, retErrors
	}

	// Check contents of contract expression
	if ast.Next != nil {
		if val, ok := ast.Next.Here.(*ast1.TySymbol); ok {
			retType.Symbol = val.Symbol
			retType.Line   = ast.GetLine()

			// Check internal definitions

			return &retType, retErrors
		}else{
			retErrors = append(retErrors, errors.New("Struct has no valid name."))
		}
	}else{
		retErrors = append(retErrors, errors.New("Struct header has no contents."))
	}

	return nil, retErrors
}




func tryPublic(ast *ast1.AST) (*ast2.VarNode, []error) {

	var retErrors []error
	var retVar    ast2.VarNode

	if !checkGenericNode(ast, "defpub") {
		retErrors = append(retErrors, utils.LineError(ast.GetLine(), "Improperly defined variable"))
		return nil, retErrors
	}

	// Check contents of contract expression
	if ast.Next != nil {
		if val, ok := ast.Next.Here.(*ast1.FnSymbol); ok {
			retVar.Symbol = val.Symbol
			retVar.Line   = ast.GetLine()

			// Check internal definitions

			return &retVar, retErrors
		}else{
			retErrors = append(retErrors, errors.New("Variable has no valid name."))
		}
	}else{
		retErrors = append(retErrors, errors.New("Variable header has no contents."))
	}

	return nil, retErrors
}
