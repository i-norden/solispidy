package parser

import (
	"strings"
	"errors"
)

// recognize unclosed paranthesis
// recognize strings


type Line struct {
	Text   string
	Tokens []string
	Number int
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
			Text: line,
			Number: i,
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

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func delete_empty (s []string) []string {
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
