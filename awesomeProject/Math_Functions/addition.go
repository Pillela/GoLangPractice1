package Math_Functions

func Addition(a int, b int) int {
	return a+b
}

func Addition2(args ...int) int{
	total:= 0
	for _,v:= range args{
		total += v
	}
	return total
}