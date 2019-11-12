package main

import (
	// "strings"
	"GoPractice/Strings/stringfuncs"
)

func main() {

	// // below is an example for raw string literal
	// fmt.Println(`hello \n abishek`)

	// //below is an example for interpreted string literals, which considers "\n" as the next line
	// fmt.Println("hello \nabishek")

	// //func exported from comparefung.go
	// //fmt.Println(stringfuncs.Comparefunc("abs", "abs"))

	// // func exported from reverseofstring.go
	// stringfuncs.Reverseofstring("testing this reverse of string func!")

	// stringfuncs.Anotherreverseofstring("Abishek")

	// //All predefined "strings" functions
	// //strings.Compare()
	// //strings.compare returns 0 when str1 and str2 are equal, returns 1 when string is greater/lengthier than str2, returns -1 when str2 is greater/lengthier than str1.
	// fmt.Println(strings.Compare("abishe", "abishek"))
	// fmt.Println(strings.Compare("Japan", "Australia")) //J>A
	// fmt.Println(strings.Compare("", " "))              // Space is less

	// fmt.Println(strings.Compare("Germany", "Germany")) // G == G
	// fmt.Println(strings.Compare("Germany", "GERMANY")) // GERMANY > Germany

	// // strings.Contains(a string, b substring) bool
	// fmt.Println(strings.Contains("Australia", "Aus"))
	// fmt.Println(strings.Contains("", "abishe "))
	// fmt.Println(strings.Contains("", ""))
	// fmt.Println(strings.Contains(" ", " "))
	// fmt.Println(strings.Contains("Japan", "JAPAN")) //case sensitive
	// fmt.Println(strings.Contains("Australia", "Australian"))

	// // strings.ContainsAny(s, char string) bool
	// fmt.Println(strings.ContainsAny("Japan", "j"))
	// fmt.Println(strings.ContainsAny("Australia", "a & r"))
	// //  Contains vs ContainsAny
	// fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) // true
	// fmt.Println(strings.Contains("Shell-12541", "1-2"))    // false

	// // strings.Count(s, sep string) int
	// fmt.Println(strings.Count("Japanese", "Japan"))
	// fmt.Println(strings.Count("Japanese", "a"))
	// fmt.Println(strings.Count("Shell-25152", "-21"))
	// fmt.Println(strings.Count("test", "")) // returns length of string +1
	// fmt.Println(strings.Count("test", " "))

	// // strings.EqualFold(s, t string) bool (not case sensitive)
	// fmt.Println(strings.EqualFold("australia", "AUSTRALIA")) //returns true since its not case sensitive
	// fmt.Println(strings.EqualFold("Australia", "ausTRALIA"))
	// fmt.Println(strings.EqualFold("Australia", "Australia & Japan"))
	// fmt.Println(strings.EqualFold(" ", " "))
	// fmt.Println(strings.EqualFold(" ", "  "))

	// //strings.Fields(s string) []string
	// testString := "8568 warren pkywy, frisco, texas 75034"
	// testArray := strings.Fields(testString)
	// fmt.Println(testArray)

	// // for _, v := range testArray {
	// // 	fmt.Println(v)
	// // }

	// string2 := ""
	// for _, v := range testArray {
	// 	string2 = string2 + " " + v
	// }
	// fmt.Println(string2)

	// string1 := "abcd"

	// for i := 0; i < len(string1); i++ {
	// 	fmt.Println(string(string1[i]))
	// }

	//stringfuncs.Removedupechars("grammar is grammar")

	stringfuncs.Anotherreverseofstring("abishek")

}
