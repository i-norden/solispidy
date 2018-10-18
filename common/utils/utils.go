package utils

import "encoding/json"

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
