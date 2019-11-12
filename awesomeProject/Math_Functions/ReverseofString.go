package Math_Functions

import "fmt"

func ReverOfString(string1 string) {
	string2 :=""
	runearray := []rune(string1)

	for i:=len(string1)-1;i>=0;i--{
		string2 = string2 + string(runearray[i])
	}
	fmt.Println(string2)
}


