package sortingfuncs

import "fmt"

//Selectionsort is exported and used main func under the sorting folder.
func Selectionsort(array1 [10]float64) [10]float64 {
	var array2 [10]float64
	replacingmaxval := 10000000.0
	fmt.Println("Min value should be:", array1[0])
	minval := array1[0]
	minvalindex := 0
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if array1[i] < minval {
				minval = array1[i]
				minvalindex = i
			}
		}
		array1[minvalindex] = replacingmaxval
		array2[j] = minval
		minval = array1[0]
	}
	return array2
}

// {23, 45, 12, 34, 45, 56, 67, 88.9, 55.5, 34.9}
