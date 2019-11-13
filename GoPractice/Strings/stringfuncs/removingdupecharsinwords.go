package stringfuncs

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

func Sorted(char []rune) []rune { //application
	k := make(map[rune]bool)
	l := []rune{}
	for _, e := range char {
		if _, v := k[e]; !v {
			k[e] = true
			l = append(l, e)
		}
	}
	return l
}
