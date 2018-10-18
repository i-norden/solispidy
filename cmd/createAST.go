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
	"log"

	"github.com/spf13/cobra"

	"github.com/i-norden/solispidy/parser"
	"github.com/i-norden/solispidy/types"
)

// createASTCmd represents the createAST command
var createASTCmd = &cobra.Command{
	Use:   "createAST",
	Short: "Creates an abstract syntax tree from the input lisp file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		createAST()
	},
}

func init() {
	rootCmd.AddCommand(createASTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createASTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createASTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createAST() {
	for _, text := range sourceFiles {
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
