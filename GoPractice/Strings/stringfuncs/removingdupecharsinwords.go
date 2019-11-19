package stringfuncs

import (
	"GoPractice/Strings/stringfuncs"
	"fmt"
	"strings"
)

// //Removedupechars is called in main.go
// func Removedupechars(s string) {
// 	stringarray := strings.Fields(s)
// 	newstring := ""
// 	word := ""
// 	for _, v := range stringarray {
// 		for i := 0; i < len(v); i++ {
// 			for j := 0; j <= len(word); j++ {
// 				if string(v[i]) != string(word[j]) {
// 					word += string(v[i])
// 				}
// 			}
// 		}
// 		newstring += word
// 	}
// 	fmt.Println(newstring)
// }
//Removedupechars export
func Removedupechars() {
	var result []string
	st := "applicationp    required knowledge"
	words := strings.Split(st, " ")
	for _, v := range words {

		char := []rune(v)
		r := string(stringfuncs.Sorted(char))
		result = append(result, r)
	}
	fmt.Println(result)
	fmt.Println(strings.Join(result, " "))
}

//Sorted export
func Sorted(char []rune) []rune { //application
	k := make(map[rune]bool)
	l := []rune{}
	for _, e := range char { //application
		if _, v := k[e]; !v {
			k[e] = true
			l = append(l, e)
		}
	}
	return l
}
