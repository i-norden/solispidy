package parser

import "strings"

type Line struct {
	Text   string
	Tokens []string
	Number int
}

type Lines []Line

func Tokenize(program string) Lines {
	var allLines Lines
	var linesOfInterest Lines

	lines := strings.Split(program, "\n")
	for i, line := range lines {

		allLines[i] = Line{
			Text: line,
			Number: i,
		}

		if string(line[0]) != ";" {
			linesOfInterest = append(linesOfInterest, allLines[i])
		}
	}

	for _, line := range linesOfInterest {
		line.Tokens = strings.FieldsFunc(line.Text, split)
	}

	return linesOfInterest
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func split(r rune) bool {
	return r == ' ' || r == '(' || r == ')'
}