package stringfuncs

import "fmt"

//Reverseofstring may be exported and used in main func under Strings folder.
func Reverseofstring(a string) {

	runearray := []rune(a)
	string2 := ""
	for i := len(runearray) - 1; i >= 0; i-- {
		string2 = string2 + string(runearray[i])
	}
	fmt.Println(string2)

}

//Anotherreverseofstring implementings swapping of rune literals.
func Anotherreverseofstring(a string) {
	runearray := []rune(a)
	for i := 0; i < len(runearray)/2; i++ { //grammar
		char1 := runearray[i]
		for j := (len(runearray) - (i + 1)); j >= 0; j-- {
			char2 := runearray[j]
			runearray[i] = char2
			runearray[j] = char1
			break
		}
	}
	fmt.Println(string(runearray))
}
