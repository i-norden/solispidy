package parser

import (
	"errors"
	"github.com/i-norden/solispidy/types"
)

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



func checkGenericNode(ast *types.AST, fnsym string) bool {
	if ast.Here == nil {
		return false
	}
	here := *ast.Here
	if val, ok := here.(types.FnSymbol); ok {
		if val.Symbol == fnsym {
			return true
		}
	}

	return false
}

func pullFnSymbol(ast *types.AST) (string, bool) {
	here := *ast.Next.Here
	if val, ok := here.(types.FnSymbol); ok {
		return val.Symbol, ok
	} else {
		return "", false
	}
}






func CheckFile(asts []types.AST) ([]types.ContractNode, []error){
	var retContracts []types.ContractNode
	var retErrors    []error

	for _, ast := range asts {
		if checkGenericNode(&ast, "def-contract") {
			contract, errs := tryContract(&ast)
			if contract != nil {
				retContracts = append(retContracts, *contract)
			}
			if errs != nil {
				retErrors    = append(retErrors, errs...)
			}
		}else{
			retErrors = append(retErrors, errors.New("Improperly defined contract"))
		}
	}

	if len(retErrors) != 0 {
		retContracts = make([]types.ContractNode, 0)
	}
	return retContracts, retErrors
}


func tryContract(ast *types.AST) (*types.ContractNode, []error){

	var retErrors []error

	// Check contents of contract expression

	return nil, retErrors
}
