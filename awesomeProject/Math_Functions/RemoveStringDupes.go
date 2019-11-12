package Math_Functions

import "bytes"

//import (
//	"bytes"
//)

func Removestringdupes(s string) {
	//runearr1 := []rune(s)
	//new_string :=""
	//b:=""
	//var val1 int=0
	//var val2 int=0
	//space:=0
	//previousspace:=0
	//for j:=0;j<len(runearr1);j++{
	//	if string(runearr1[j])==" "{
	//		val2=j
	//		for i:=val1;i<val2;i++{
	//			for k:=i+1;k<=val2;k++{
	//				if string(runearr1[i])== string(runearr1[k]){
	//				}else{
	//					b = string(runearr1[i])
	//				}
	//			}
	//			new_string = new_string+b
	//		}
	//		val1=val2
	//		//fmt.Println(val1)
	//	}
	//}
	//return new_string

	//fmt.Println(strings.Fields(s)[0])
	//array1 := strings.Fields(s)
	//array2 := make([]string, length1)
	//string1:=s
	//grammar
	//string2:=""
	//rune1:= []rune(s)
	//b:=""

	//for i:=0;i<=len(rune1);i++{
	//	for _,v:=range &string2{
	//		if rune1[i] != rune(v){
	//			*string2 = *string2+ string(rune1[i])
	//		}
	//	}
	//}
	//fmt.Println("test")
	//fmt.Println(string2)

	//for i:=0;i<len(rune1);i++{
	//	for j:=i+1;j<len(rune1);j++{
	//		if rune1[i]!=rune1[j]{
	//			b=string(rune1[i])
	//		}else{
	//			continue
	//		}
	//	}
	//	string2 = string2 + b
	//}
	//fmt.Println(string2)


}

func removeDups(s string) string {
	var buf bytes.Buffer
	var last rune
	for i, r := range s {
		if r != last || i == 0 {
			buf.WriteRune(r)
			last = r
		}
	}
	return buf.String()
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}