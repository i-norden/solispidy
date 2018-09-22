package main

import "io/ioutil"
import "os"
import "fmt"






func loadSourceFiles (files []string) ([]string, error) {
  numfiles := len(files)
  var filetexts []string
  for i := 0; i < numfiles; i++ {
    text, err:= ioutil.ReadFile(files[i])
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

  _, errs := loadSourceFiles(files)

  if(errs != nil){
    fmt.Println(errs)
  }else{
    fmt.Println("Files loaded successfully.")
  }
}
