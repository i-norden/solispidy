package main

import "io/ioutil"
import "os"
import "fmt"
import "../parser"






func loadSourceFiles (files []string) ([]string, error) {
  var filetexts []string
  for _, file := range files {
    text, err:= ioutil.ReadFile(file)
    if(err != nil){
      out := []string{}
      return out, err
    }
    filetexts = append(filetexts, string(text))
  }
  return filetexts, nil
}







func main(){
  args := os.Args[1:]

  // We need to filter this later. Some parameters may be parameters, for
  // example, perhaps the path to solidity if it can't be found in the expected
  // locations.
  files := args

  fmt.Println(files)

  texts, errs := loadSourceFiles(files)

  if(errs != nil){
    fmt.Println(errs)
    return
  }else{
    fmt.Println("Files loaded successfully.")
  }


  for _, text := range texts {
    lines, errs := parser.Tokenize(text)
    if(errs != nil){
      fmt.Println(errs)
      return
    }else{
      for _, line := range lines {
        fmt.Println(line)
      }
    }
  }

}
