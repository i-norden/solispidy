package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/i-norden/solispidy/parser"
	"github.com/i-norden/solispidy/types"
)

func loadSourceFiles(files []string) ([]string, error) {
	var filetexts []string
	for _, file := range files {
		text, err := ioutil.ReadFile(file)
		if err != nil {
			out := []string{}
			return out, err
		}
		filetexts = append(filetexts, string(text))
	}
	return filetexts, nil
}

func main() {
	args := os.Args[1:]

	// We need to filter this later. Some parameters may be parameters, for
	// example, perhaps the path to solidity if it can't be found in the expected
	// locations.
	files := args

	fmt.Println(files)

	texts, err := loadSourceFiles(files)

	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Files loaded successfully.")
	}

	for _, text := range texts {
		lines, err := parser.Tokenize(text)
		if err != nil {
			log.Fatal(err)
			return
		} else {
			tokens, err := parser.ReadFromLines(lines)
			if err != nil {
				log.Fatal(err)
			}
			ast, err := parser.MakeAST(tokens, types.AST{}, 0)
			if err != nil {
				log.Fatal(err)
			}
			str := parser.PrettyPrint(ast)
			println(str)
		}
	}

}
