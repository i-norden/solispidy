package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)





func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func Split(r rune) bool {
	return r == ' ' || r == '(' || r == ')'
}



func LineError(l int64, err string) error{
	errhead := fmt.Sprintf("On Line %i:\n  ", l)
	// The replace is there for indentation; it shifts the error message over by 2 spaces.
	return fmt.Errorf(errhead, strings.Replace(err, "\n", "  \n", -1))
}
