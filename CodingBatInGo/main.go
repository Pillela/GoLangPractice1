package main

import (
	"CodingBatInGo/All_programs"
	"fmt"
)

var bool_var bool

func main(){
	week_day := true
	vaca := false
	bool_var := All_programs.Sleep_in(week_day, vaca)
	fmt.Println(bool_var)

}
