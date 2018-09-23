package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/i-norden/solispidy/types"
	"os"
	"strconv"
	"strings"
)

// recognize unclosed paranthesis
// recognize strings
// identify any capitalized word as type

type Line struct {
	Text   string
	Tokens []string
	Number int64
}

type Lines []*Line

func Tokenize(program string) (Lines, error) {

	leftPars := strings.Count(program, "(")
	rightPars := strings.Count(program, ")")
	if rightPars > leftPars {
		return nil, errors.New("Missing opening parenthesis")
	}
	if leftPars > rightPars {
		return nil, errors.New("Missing closing parenthesis")
	}

	var linesOfInterest Lines

	lines := strings.Split(program, "\n")

	for i, line := range lines {

		l := Line{
			Text:   line,
			Number: int64(i),
		}

		if (line != "") && (string(line[0]) != ";") {
			linesOfInterest = append(linesOfInterest, &l)
		}
	}

	for _, line := range linesOfInterest {
		var tempStr string
		tempStr = strings.Replace(line.Text, "(", "( ", -1)
		tempStr = strings.Replace(tempStr, ")", " )", -1)
		line.Tokens = strings.Split(tempStr, " ")
		line.Tokens = delete_empty(line.Tokens)
	}

	return linesOfInterest, nil
}

func ReadFromLines(lines Lines) ([]types.Symbol, error) {
	var tokens []types.Symbol

	for _, line := range lines {
		parsedTokens, err := ReadFromTokens(line.Tokens, line.Number)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, parsedTokens...)
	}

	return tokens, nil
}

func ReadFromTokens(tokens []string, ln int64) ([]types.Symbol, error) {
	if len(tokens) == 0 {
		return nil, errors.New("Unexpected EOF")
	}

	result := make([]types.Symbol, len(tokens))

	// Need to put spaces in front and/or behind quotations so that you can find em
	// try to make this work to sort out quotes/strings
	/*
		quo1 := find(tokens, `"`)
		if quo1 != -1 {
			quo2 := find(tokens[quo1+1:], `"`)
			join := strings.Join(tokens[quo1+1:quo2], "")
			new := append(tokens[quo1+1:], join)
			tokens = append(new, tokens[:quo2]...)
		}
	*/

	for i, token := range tokens {
		result[i] = atom(token, ln)
	}

	return result, nil
}

func MakeAST(symbols []types.Symbol, ast types.AST, count int) (*types.AST, error) {
	symbol := symbols[0]
	symbols = symbols[1:]

	switch t := symbol.(type) {
	case types.LeftPar:
		count += 1
		return MakeAST(symbols, ast, count)
	case types.RightPar:
		count -= 1
		if count == 0 {
			ast.Next = nil
			return &ast, nil
		}
		return MakeAST(symbols, ast, count)
	case types.AssertNode:
		fmt.Fprintf(os.Stderr, "Type: %v\r\n", t)
		return nil, nil
	default:
		ast.Here = &symbol
		next, err := MakeAST(symbols, types.AST{}, count)
		if err != nil {
			return nil, err
		}
		ast.Next = next
		return &ast, nil
	}
}

func atom(token string, ln int64) types.Symbol {
	if token == "(" {
		return types.LeftPar{LPId: 1, Line: ln}
	}
	if token == ")" {
		return types.RightPar{RPId: 1, Line: ln}
	}
	if token == "False" || token == "false" {
		return types.CnstBool{Data: false, Line: ln}
	}
	if token == "True" || token == "true" {
		return types.CnstBool{Data: true, Line: ln}
	}

	i, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return types.CnstStr{Data: token, Line: ln}
	}
	ui := uint64(i)
	return types.CnstInt{Data: [4]uint64{ui}, Line: ln}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func split(r rune) bool {
	return r == ' ' || r == '(' || r == ')'
}
