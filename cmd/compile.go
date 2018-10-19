// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"errors"

	"github.com/i-norden/solispidy/parser"
	symbolizer "github.com/i-norden/solispidy/symbolizer"

	ast1 "github.com/i-norden/solispidy/parser/types"
	//ast2 "github.com/i-norden/solispidy/symbolizer/types"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles input file into solidity",
	Long: `Parses the input file into an AST that is then used
to undergo formal verification of the input code. If the code
passes formal verification it is optimized for gas consumption
and compiled down to solidity for contract publication.`,
	Run: func(cmd *cobra.Command, args []string) {
		compile()
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func compile() {

	var asts []ast1.AST

	for _, text := range sourceFiles {

		p := new(parser.Parser)

		err := p.Parse(text)
		if err != nil {
			log.Fatal(err)
		}

		asts = append(asts, *p.Ast)
	}

	_, errs := symbolizer.CheckFile(asts)

	if errs != nil {
		for _, err := range errs {
			fmt.Println(err)
		}
		log.Fatal(errors.New("Unable to compile due to above errors."))
	}else{
		fmt.Println("Success")
	}

}
