package Math_Functions

import "fmt"

func Array_Max_Value(){
	var length int
	fmt.Print("Please enter the length of an array")
	fmt.Scanf("%d",&length)
	fmt.Println("Please enter values of the array")
	//var i int
	//using a slice as we cant do this using an array
	array1 := make([]int, length, length)
	//trie
	//array1:= new([length]int)
	for i:=0; i<length; i++{
		fmt.Scanln(&array1[i])
	}
	fmt.Println(array1)

	var max =0
	var x int
	for x=0; x<length; x++{
		if max<array1[x]{
			max=array1[x]
		}else{
			max=max
		}
	}
	fmt.Println("Max value within given array is:", max)

}
