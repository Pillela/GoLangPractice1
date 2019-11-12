package Math_Functions

import (
	"math/rand"
	"time"
)

func Random_Number(x [30]int) [30]int{

	s1:= rand.NewSource(time.Now().UnixNano())
	r1:= rand.New(s1)

	for i:=0;i<len(x);i++{
		x[i]=r1.Intn(30)
	}
	return x
}