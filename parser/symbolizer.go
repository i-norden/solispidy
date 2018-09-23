package parser


import "../types"
import "errors"



// cdr is an old lisp function that returns the "next" value in a list
// cdrlist returns an array of the next val, next of that, etc.
func cdrlist(ast *types.AST) ([]*types.AST){
  i := ast.Next;
  var ret []*types.AST
  for i != nil {
    ret = append(ret, i)
    i = i.Next
  }
  return ret
}



func cdarlist(ast *types.AST) ([]*interface{}){
  i := ast.Next;
  var ret []*interface{}
  for i != nil {
    ret = append(ret, i.Here)
    i = i.Next
  }
  return ret
}


func nodelist(ast *types.AST) ([]*types.AST){
  i := ast
  var ret []*types.AST
  for i != nil {
    ret = append(ret, i)
    i = i.Next
  }
  return ret
}



func carlist(ast *types.AST) ([]*interface{}){
  i := ast
  var ret []*interface{}
  for i != nil {
    ret = append(ret, i.Here)
    i = i.Next
  }
  return ret
}






/*
func checkContractDef(ast *types.AST) (error){
  header  := *ast.Here
  val, ok := header.(types.FnSymbol)
  if !ok || (val.Symbol != "def-contract") {
    return errors.New("Expected (def-contract ...)")
  }else{
    return nil
  }

  defs := cdarlist(ast)
  for _, v := range defs {
    derefv := *v
    val, ok := derefv.(types.AST)
    if ok {
      var check [5]error
      check[0] = checkFnDef(&val)
      check[1] = checkTyDef(&val)
      check[2] = checkPubvarDef(&val)
      check[3] = checkPrifunDef(&val)
      check[4] = checkAssert(&val)

      for i:=0; i<5; i++ {
        if check[i] != nil {
          return nil
        }
      }
      return errors.New("Contract features invalid definition")

    }else{
      return errors.New("Expected a definition")
    }
  }
  return nil
}


func checkFnDef(ast *types.AST) (error){
  header  := *ast.Here
  val, ok := header.(types.FnSymbol)
  if !ok || (val.Symbol != "defn") {
    return errors.New("Expected (defn ...)")
  }else{
    return nil
  }

  //_ := cdarlist(ast)
  return nil
}


func checkTyDef(ast *types.AST) (error){
  header  := *ast.Here
  val, ok := header.(types.FnSymbol)
  if !ok || (val.Symbol != "defty") {
    return errors.New("Expected (defty ...)")
  }else{
    return nil
  }

  //_ := cdarlist(ast)
  return nil
}


func checkPubvarDef(ast *types.AST) (error){
  header  := *ast.Here
  val, ok := header.(types.FnSymbol)
  if !ok || (val.Symbol != "def-pubvar") {
    return errors.New("Expected (def-pubvar ...)")
  }else{
    return nil
  }

  //_ := cdarlist(ast)
  return nil
}


func checkPrifunDef(ast *types.AST) (error){
  header  := *ast.Here
  val, ok := header.(types.FnSymbol)
  if !ok || (val.Symbol != "def-prifun") {
    return errors.New("Expected (def-prifun ...)")
  }else{
    return nil
  }

  //_ := cdarlist(ast)
  return nil
}


func checkAssert(ast *types.AST) (error){
  header  := *ast.Here
  val, ok := header.(types.FnSymbol)
  if !ok || (val.Symbol != "assert-state") {
    return errors.New("Expected (assert-state ...)")
  }else{
    return nil
  }

  //defs := cdarlist(ast)
  return nil
}


func checkPar(ast *types.AST) (error){
  if ast.Next == nil {
    return errors.New(fmt.Sprintf("L:%d : Expected a second value in parameter list."))
  }
  head := *ast.Here
  tail := *ast.Next.Here

  if _, ok := head.(types.FnSymbol); ok {
    if _, ok := tail.(types.TySymbol); ok {
      return nil
    }
  }
  return errors.New(fmt.Sprintf("L:%d : Improperly formatted parameter."))
}


func checkVar(ast *types.AST) (error){
  if ast.Next == nil {
    return errors.New(fmt.Sprintf("L:%d : Expected a 3 values in variable definition."))
  }
  if ast.Next.Next == nil{
    return errors.New(fmt.Sprintf("L:%d : Expected a 3 values in variable definition."))
  }
  head := *ast.Here
  ty   := *ast.Next.Here
  tail := *ast.Next.Next.Here

  if _, ok := head.(types.FnSymbol); ok {
    if _, ok := ty.(types.TySymbol); ok {
      // Add a checkExpr here later
      return nil
    }
  }
  return errors.New(fmt.Sprintf("L:%d : Improperly formatted parameter."))
}

/*
func checkExpr(ast *types.AST) (error){

}*/


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



func pullFnSymbol(ast *types.AST) (string, bool){
  here := *ast.Next.Here
  if val, ok := here.(types.FnSymbol); ok{
    return val.Symbol, ok
  }else{
    return "", false
  }
}



func checkContractDef(ast *types.AST) (*types.ContractNode, error){
  if val, err := checkGenericNode(ast, "def-contract"); err != nil {
    here := *val.Here
    here_ := here.(types.Symbol)
    ret := types.ContractNode{(here_.GetLine()), []types.FnNode{}, []types.TyNode{}, []types.VarNode{}, []types.AssertNode{}, []types.FieldNode{}, types.EmptyTable(), ""}
    sym0, _ := pullFnSymbol(ast.Next)
    ret.Symbol = sym0

    // Iterate over defs, typeswitch on defs and add them to contract arrays.
    // If anything defaults, return an error.
    //cdrlist()

  }
  return nil, errors.New("Failed to parse contract definition")
}


func checkFnDef(ast *types.AST) (*types.FnNode, error){
  if val, err := checkGenericNode(ast, "defn"); err != nil {
    here := *val.Here
    here_ := here.(types.Symbol)
    ret := types.FnNode{(here_.GetLine()), "", []types.TypeNote{}, []types.TypeNote{}, true, nil, types.EmptyTable()}
    sym0, _ := pullFnSymbol(ast.Next)
    ret.Symbol = sym0

    // Check if inpars/expars are valid

    // Check if definition is valid

    // Insert parnames into Symbol Table

  }
  return nil, errors.New("Failed to parse function definition")
}


func checkTyDef(ast *types.AST) (*types.TyNode, error){
  if val, err := checkGenericNode(ast, "defty"); err != nil {
    here := *val.Here
    here_ := here.(types.Symbol)
    ret := types.TyNode{(here_.GetLine()), "", []types.PairNode{}}
    sym0, _ := pullFnSymbol(ast.Next)
    ret.Symbol = sym0

    // Check if fields are valid

  }
  return nil, errors.New("Failed to parse type definition")
}
