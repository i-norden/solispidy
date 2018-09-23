package parser

import (
	"errors"
	"github.com/i-norden/solispidy/types"
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

	var result []types.Symbol
	token := tokens[0]
	copy(tokens, tokens[1:])
	tokens = tokens[:len(tokens)-1]

	if token == "(" {
		for tokens[0] != ")" {
			n, err := ReadFromTokens(tokens, ln)
			if err != nil {
				return nil, err
			}
			result = append(result, n...)
		}
		copy(tokens, tokens[1:])
		tokens = tokens[:len(tokens)-1]

		return result, nil
	}
	if token == ")" {
		return nil, errors.New("Unexpected )")
	}

	result = append(result, atom(token, ln))

	return result, nil
}

func atom(token string, ln int64) types.Symbol {
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
