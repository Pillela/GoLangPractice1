package Math_Functions

import "fmt"

	type Intnums struct {
		A int
		B int
	}

	type Floatnums struct{
		A float64
		B float64
	}

	func (a Intnums) Method1() float64{
		return float64(a.A+a.B)
	}

	func (a Floatnums) Method1() float64{
		return a.A+a.B
	}

	func (a Floatnums) Method2() float64{
		return a.A*a.B
	}
	type Sum interface{
		Method1() float64
		Method2() float64
	}

	func Total_addition(a ...Sum){
		//fmt.Println(a)

		for _, x := range a{
			fmt.Println(x.Method1())
			fmt.Println(x.Method2())
		}
	}

