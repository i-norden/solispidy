package parser

import (
	"errors"
	"github.com/i-norden/solispidy/common/utils"
	"github.com/i-norden/solispidy/parser/types"
	"strconv"
	"strings"
)

// recognize unclosed paranthesis
// recognize strings
// identify any capitalized word as type

// Types to represent program line-by-line
type Line struct {
	Text   string
	Tokens []string
	Number int64
}

type Lines []*Line

// Parser struct to hold symbols and resulting ast as
// well as current index within the symbols array
type Parser struct {
	Symbols []types.Symbol
	Ast     *types.AST
	index   int
}

// Exported parser method that sanitizes and creates
func (p *Parser) Parse(program string) (err error) {
	lines, err := createLines(program)
	if err != nil {
		return err
	}

	tokens, err := readFromLines(lines)
	if err != nil {
		return err
	}

	p.Symbols = tokens
	p.Ast = &types.AST{}
	p.index = 0

	p.Ast, err = p.makeAST(types.AST{}, 0)
	if err != nil {
		p.Ast = nil
		return err
	}

	return nil
}

// Internal parser method used recursively to parse a program
func (p *Parser) makeAST(ast types.AST, count int) (*types.AST, error) {
	symbol := p.Symbols[p.index]
	p.index++

	switch symbol.(type) {
	case *types.LeftPar:
		count += 1
		if count > 1 {
			count = 1
			here, err := p.makeAST(types.AST{}, count)
			if err != nil {
				return nil, err
			}
			ast.Here = here

			next, err := p.makeAST(types.AST{}, count)
			if err != nil {
				return nil, err
			}
			ast.Next = next

			return &ast, nil
		}

		return p.makeAST(ast, count)
	case *types.RightPar:
		count -= 1
		if count == 0 {
			ast.Next = nil
			return &ast, nil
		}

		return p.makeAST(ast, count)
	default:
		ast.Here = symbol

		next, err := p.makeAST(types.AST{}, count)
		if err != nil {
			return nil, err
		}
		ast.Next = next

		return &ast, nil
	}
}

// Helper functions

// Function to cast token to matching symbol
func atom(token string, ln int64) types.Symbol {
	if token == "(" {
		return &types.LeftPar{LPId: 1, Line: ln}
	}
	if token == ")" {
		return &types.RightPar{RPId: 1, Line: ln}
	}
	if token == "False" || token == "false" {
		return &types.CnstBool{Data: false, Line: ln}
	}
	if token == "True" || token == "true" {
		return &types.CnstBool{Data: true, Line: ln}
	}

	i, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return &types.CnstStr{Data: token, Line: ln}
	}
	ui := uint64(i)
	return &types.CnstInt{Data: [4]uint64{ui}, Line: ln}
}

// Functions to read program line-by-line and
// and convert raw tokens to typed symbols
func readFromLines(lines Lines) ([]types.Symbol, error) {
	var tokens []types.Symbol

	for _, line := range lines {
		symbols, err := readFromTokens(line.Tokens, line.Number)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, symbols...)
	}

	return tokens, nil
}

func readFromTokens(tokens []string, ln int64) ([]types.Symbol, error) {
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

// Functions to sanitize program and cut
// it up into a slice of lines
func createLines(program string) (linesOfInterest Lines, err error) {

	linesOfInterest, err = sanitize(program)

	for _, line := range linesOfInterest {
		var tempStr string
		tempStr = strings.Replace(line.Text, "(", "( ", -1)
		tempStr = strings.Replace(tempStr, ")", " )", -1)
		line.Tokens = strings.Split(tempStr, " ")
		line.Tokens = utils.DeleteEmpty(line.Tokens)
	}

	return linesOfInterest, nil
}

func sanitize(program string) (linesOfInterest Lines, err error) {

	leftPars := strings.Count(program, "(")
	rightPars := strings.Count(program, ")")
	if rightPars > leftPars {
		return nil, errors.New("Missing opening parenthesis")
	}
	if leftPars > rightPars {
		return nil, errors.New("Missing closing parenthesis")
	}

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

	return
}
