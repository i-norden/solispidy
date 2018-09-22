package main










type AST struct{
  here *interface{}
  next *interface{}
}






type FnId struct{
  fnid int64
}






type TyId struct{
  tyid int64
}





type AssertId struct{
  assertid int64
}





type CnstInt struct{
  data [4]uint64
}






type CnstStr struct{
  data string
}





type CnstBool struct{
  data bool
}





type SymbolTable struct{
  fndefs     []AST
  tydefs     []AST
  assertdefs []AST
}
