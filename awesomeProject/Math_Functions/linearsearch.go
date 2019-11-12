package Math_Functions

func LinearSearch(a []int, x int) int{
	for i:=0;i<len(a);i++{
		if a[i]==x{
			return i
		}
	}
	return -1
}
