package parser

import (
	"errors"
	//"../types"
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

func checkContractDef(ast *types.AST) (*types.ContractNode, error) {
	if val, err := checkGenericNode(ast, "def-contract"); err != nil {
		here := *val.Here
		here_ := here.(types.Symbol)
		ret := types.ContractNode{(here_.GetLine()), []types.FnNode{}, []types.TyNode{}, []types.VarNode{}, []types.AssertNode{}, []types.FieldNode{}, types.EmptyTable(), ""}
		sym0, _ := pullFnSymbol(ast.Next)
		ret.Symbol = sym0

		// Iterate over defs, typeswitch on defs and add them to contract arrays.
		// If anything defaults, return an error.
		//cdrlist()

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
