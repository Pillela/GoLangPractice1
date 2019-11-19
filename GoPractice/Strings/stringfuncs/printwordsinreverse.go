package stringfuncs

import (
	"strings"
)

//Printwordsinreverse export
func Printwordsinreverse(s string) string {
	wordarray := strings.Fields(s)
	string2 := ""
	for i := len(wordarray) - 1; i >= 0; i-- {
		string2 = string2 + " " + wordarray[i]
	}
	return string2
}
