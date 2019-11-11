package main

import (
	"GoPractice/Sorting/sortingfuncs"
	"fmt"
)

func main() {
	array1 := [10]float64{23, 45, 12, 34, 45, 56, 67, 88.9, 55.5, 34.9}
	fmt.Println(len(array1))
	array2 := sortingfuncs.Selectionsort(array1)
	fmt.Println(array2)

}
