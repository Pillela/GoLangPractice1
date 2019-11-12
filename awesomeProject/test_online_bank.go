package main

import (
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
	//"fmt"
	// "io"
	// " math"
	// "net"
	// "net/http"
	// "net/rpc"
	// "sync"
	// "time"
)

var x int = 10

func f() {
	fmt.Println(x)
}

func add(x int, y int) (int, int, int) {
	return (x + y), (x - y), (x * y)
}

func panik() {
	panic("Somethings messed up bruh!")
}

func car_stuff(a Car) {
	fmt.Println()
}

//Car struct
type Car struct {
	make      string
	car_model string
	year      int
	price     float64
}

func naked_return(x int) (z, y int) {
	z = x + 1
	y = x + 2
	return
}

//const (a = 1 << 50
//b = a >> 49
//)
//var b1 = "Krishna"
func main() {

	fmt.Println("hello there!")
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println("This is go \npractice")
	////c := "Hare " + b1
	//fmt.Printf("Type of b1 = %T",b1)
	//var number = 1
	//var string1 = "1" + fmt.Sprint(number)
	//fmt.Println("\n")
	//fmt.Println(string1)
	////fmt.Println("\n I'm the",string(number), "yeah!")

	//var var1 float64
	//fmt.Println("Enter var1:")
	//fmt.Scanf("%f",&var1)
	//var var2  = 2.345
	//err:= fmt.Println(var1 * var2)
	//fmt.Println(err)

	//var i int
	//fmt.Println("Which math table would you like to see?:")
	//fmt.Scanf("%d",&i)
	//val:=1
	//for val<=20{
	//	fmt.Println(fmt.Sprint(i),"*",fmt.Sprint(val)," = ",i*val )
	//	val++
	//}
	//var i int
	//for i=0; i<=100; i++{
	//	if i%2==0{
	//		fmt.Println(i, "is even")
	//	} else{
	//		fmt.Println(i, "is odd")
	//	}
	//}

	//var i int
	//for i=1; i<=100; i++{
	//	if i%3==0 && i%5!=0{
	//		fmt.Println("Fizz", i)
	//	}else if i%5==0 && i%3!=0{
	//		fmt.Println("Buzz", i)
	//	}else if i%5==0 && i%3==0{
	//		fmt.Println("FizzBuzz!", i)
	//	} else{
	//		fmt.Println("No Fizz, No Buzz, No FizzBuzz", i)
	//	}
	//}

	//var gender string
	//fmt.Println("Please enter the gender:")
	//fmt.Scanf("%s", &gender)
	//if gender!=""{
	//switch gender{
	//case "male": fmt.Println("Hello Handsome!")
	//case "female": fmt.Println("Hello Beautiful!")
	//default: fmt.Println("Hello There!")
	//}
	//}

	//var i[5] int
	//var x int =0
	//for x<=4{
	//i[x]=x
	//x++
	//}
	//fmt.Println(i)

	//var x string
	//x="abishek"
	//fmt.Println(x)

	//Max value in a given int array
	//Math_Functions.Array_Max_Value()

	//primes
	//fmt.Println("Enter the range of numbers,both low and high")
	//var low int
	//var high int
	//fmt.Scanf("%d",&low)
	//fmt.Scanf("%d",&high)
	//var i int
	//var j int
	//for i=low; i<=high; i++{
	//	for j=1;j<=high;j++{
	//		if i%i==0 && i%1==0 && i%j!=0{
	//			fmt.Println(i)
	//		}
	//	}
	//}

	//abc:= [3]string{"a","b","c"}
	//fmt.Println(abc)

	//names := [4]string{
	//	"Ai",
	//	"Bi",
	//	"Ci",
	//	"Di",
	//}
	//
	//a:= names[0:2]
	//b:= names[2:4]
	//
	//b[0]="abishek"
	//fmt.Println(a)
	//fmt.Println(b)
	//
	//fmt.Println(names)

	//food_menu := [6] string{"Biryani", "Pulao", "curd_rice", "lassi", "ice-cream", "kheer"}
	//var dessert_menu[] string = food_menu[3:6]
	//
	//fmt.Println(food_menu)
	//fmt.Println(dessert_menu)

	// abc := []int {1,2,3,4}
	//fmt.Println(abc)
	//
	// def :=[]string{"java","python","C","Go",}
	// fmt.Println(def)
	//
	// type manual_struct struct {
	//	a int
	//	b string
	// }
	// var i int
	// for i=0;i<4;i++{

	//slices - append and copy

	//var x[5] int
	//fmt.Println(x)
	//
	//y :=[5]int {1,2,3,4,5}
	//fmt.Println(y)
	//var total int
	//for _, value:= range y{
	//	total += value
	//}
	//
	//fmt.Println(total)
	//
	//z := y[0:3]
	//fmt.Println(z)
	//
	//z2:= append(z[:2],4,5,6,7)
	//fmt.Println(z2)
	//
	//z3:= make([]int,len(z2))
	//copy(z3,z2)
	//fmt.Println(z3)

	//array within an array
	//board:= [][]string{
	//	[]string {"_","_","_"},
	//	[]string {"_","_","_"},
	//}
	//
	//fmt.Println(board)
	//
	//for i:=0;i<len(board);i++{
	//	fmt.Println(board[i])
	//}

	//running the code within a loop when a condition is met
	//elements:= map[string]string{
	//	"H":"Hydrogen",
	//	"Li":"Lithium",
	//	"Ab":"Abishek",
	//}
	//
	//fmt.Println(elements["H"])
	//
	//name, ok:= elements["H"]
	//fmt.Println(name, ok)
	//
	//if _,ok:=elements["cd"]; !ok{
	//	fmt.Println("cd is available in elements map")
	//}

	//x := []int{1,2,3,4}

	//x:= make([]string,5)
	//x :=[5]int {1,2,3,4,5}
	//y := map[string]int{"a":1, "b":2}
	//fmt.Println(x)
	//fmt.Println(y)

	//x:= [][][]string{
	//	[][]string{"A"
	//		[]string{},
	//	},
	//	[][]string{"B"
	//		[]string{},
	//	},
	//	[][]string{"C"
	//		[]string{},
	//	},
	//}

	//fmt.Println(x)

	//POINTERS
	//a:=1453465543456576576
	//b:=2423423453453463456
	//c:= Math_Functions.Addition(a,b)
	//fmt.Println(c)
	//d := &c
	//fmt.Printf("%T",d)
	//fmt.Println("\n",d)

	//changing the value of a variable from 1 to 0
	//x:= 1
	//Math_Functions.Make_zero(&x)
	//fmt.Println(x)

	//below code is a bit confusing but its correct. when we do x:=new(int), we are allocating an address without a value and during such declarations the type of x is always *int which means its refering to an address and to get the value of x you must use *x.
	//x:= new(int)
	//Math_Functions.Make_zero(x)
	//fmt.Println(*x)

	//SLICE APPENDS

	//var s[] int
	//fmt.Println(s,len(s),cap(s))
	//
	//
	//s=append(s,0)
	//fmt.Println(s,len(s),cap(s))
	//
	//
	//s=append(s,0,1)
	//fmt.Println(s,len(s),cap(s))
	//
	//s=append(s,0,1,2)
	//fmt.Println(s,len(s),cap(s))

	//TESTING NAKED RETURN FUNCTION
	//var x int = 1
	//a,b :=  naked_return(x)
	//fmt.Println(a,b)

	//TESTING RANDOM NUMBER GENERATOR FUNCTION. This function generates the same output everytime, to see a different output we should give it a seed that changes.
	//var arr1 [30]int
	//y:= Math_Functions.Random_Number(arr1)
	//fmt.Println(y)

	//for i:=0;i<30;i++{
	//	fmt.Print(rand.Intn(30),",")
	//}

	//random number generator using seed
	//s1:= rand.NewSource(time.Now().UnixNano())
	//r1:= rand.New(s1)
	//for i:=0;i<30;i++{
	//	fmt.Print(r1.Intn(30),",")
	//}

	//using range function on the above y array generated using random number generator
	//for a,b:=range y{
	//	fmt.Println(a,b)
	//}

	//	using car struct
	//car1:= car{"toyota", "camry",2012,7500.00}
	//fmt.Println(car1)

	//maps. MAPS MUST BE INITIALIZED USING "MAKE" KEYWORD
	//var x map[string]int
	//x["name"]=1
	//fmt.Println(x)
	//above code doesnt display x values. below code does.

	//a:= make(map[string]string)
	//a["name"]="abishek"
	//a["age"]="26"
	//a["city"]="Frisco"
	//a["gender"]="male"
	//fmt.Println(a)
	//fmt.Println(a["name"])
	//delete(a, "age")
	//fmt.Println(a["age"])
	//fmt.Println(a)
	//length of a map will change as we add more items. where as in arrays size remains constant.

	//testing append function on arrays
	//x := []int {1,2,3,4,5}
	//fmt.Println(x)
	//x= append(x,6)
	//fmt.Println(x)

	// map of strings to maps of strings to strings.
	//sage_it :=map[string]map[string]string{
	//	"Resource Managers": map[string]string{
	//		"Name":"Venkat",
	//		"Age": "35",
	//	},
	//	"Sales_Managers": map[string]string{
	//		"Name":"XYZ",
	//		"Age":"40",
	//	},
	//}

	//sage_it:= make(map[string]map[string]string)
	//sage_it=["Resource_Managers"]
	//sage_it=["Resource_Managers"]
	//sage_it["Resource_Managers"]="Name"
	//sage_it["Resource_Managers"]="Age"
	//sage_it["Resource_Managers"]["Name"]="Abishek"
	//sage_it["Resource_Managers"]["Age"]="35"
	//sage_it["Sales_Managers"]["Name"]="XYZ"
	//sage_it["Resource_Managers"]["Age"]="40"
	//
	//fmt.Println(sage_it)
	//fmt.Println(sage_it["Resource Managers"])
	//fmt.Println(sage_it["Resource Managers"]["Name"])

	//plaing with struct type car
	//var car1 car
	//car1.make="toyota"
	//fmt.Println(car1)

	//car1 := new(car)
	//car1.make="toyota"
	//fmt.Println(car1)

	//car1:= car{make:"toyota", year:2012, car_model:"camry"} //or car1 := car{"toyota",2012, "camry" }
	//fmt.Println(car1)

	//car1 := Car{"toyota","camry",2012, 7500}
	//fmt.Printf("%T", car1)

	//average function with float
	//x:= [5]float64{34.5,22.6,33.8,12.7,88.43}
	//total:=0.0
	//for i:=0;i<len(x);i++{
	//	total +=x[i]
	//}
	//avg:= total/float64(len(x))
	//fmt.Println(avg)

	//testing panic function
	//panik()

	//x:=[]int{1,2,3,4,5,5,6,7,8,8,9,7,6,5,4,3,22,3,4,4,5,5,6,7,7,8,8,8,8,8,}
	//fmt.Println(Math_Functions.Addition2(x ...))

	//writing a function within a func and assiging to a variable. SO COOOOL!

	//total:=0
	//a:=[]int{1,2,3,4,5,9,6,100}
	//x:= func (a[] int) int{
	//	for _,v:= range a{
	//		total+=v
	//}
	//return total
	//}
	//
	//fmt.Println(x(a))

	//accessing f() defined in the beginning within the same package. f() uses a global variable defined within main package.
	//f()

	//testing functions (f2, f3) defined below, after the main function. It works! and the order of the functions is not mandatory.
	//fmt.Println(f2())

	//testing closure functions with and with out naked returns.
	//x:=0
	//var1:= func() int{
	//	x++
	//	return x
	//}
	//fmt.Println(var1())
	//fmt.Println(var1())

	//with naked return
	//x:=0
	//var1:= func() (a int){
	//	a = x
	//	x++
	//	return
	//}
	//fmt.Println(var1())
	//fmt.Println(var1())
	//fmt.Println(var1())

	//testing advanced closure function abc() written below.
	//var1 := abc()
	//fmt.Println(var1())
	//fmt.Println(var1())
	//fmt.Println(var1())

	//nextEven := makeEvenGenerator()
	//fmt.Println(nextEven()) // 0
	//fmt.Println(nextEven()) // 2
	//fmt.Println(nextEven()) // 4

	//var1:= abc1()
	//fmt.Println(var1())
	//fmt.Println(var1())
	//fmt.Println(var1())

	//recursion examples
	//var1:= recur1(5)
	//fmt.Println(var1)

	//testing recursion on fibonacci series
	//fmt.Println(fib(6))

	// printing output based on a condition implemented using multiple return func()
	//x, y, z:= f1(0)
	//if z==true{
	//	fmt.Println(x,y)
	//}else{
	//	fmt.Println("not much to print")
	//}

	//testing defer function
	//defer second()
	//first()

	//even odd problem.
	//fmt.Println(Math_Functions.Half_evenorodd(93))

	//i := 13
	//fmt.Println(float64(i/2))

	//var1:= 0
	//fmt.Println()

	//make string
	//car_model string
	//year int
	//price float64

	//c1:= Car{"toyota","camry",2010,7499.99}
	//car_info(&c1)
	//fmt.Println(c1.price)

	//calling car_info2 method created specifically to display car info.
	//c1.car_info2()
	//fmt.Println(c1.price)

	//var1 := car_details{c1, "sdfhjk329487sad", 2019, "yr33442"}
	//var1.print_car_details()
	//var1.Car.car_info2()
	//

	//testing employee embedded types.
	//emp_obj1 := Math_Functions.Employee_biodata{"abishek", 26, "male"}
	//emp_obj1.Basic_info()
	//fmt.Println(emp_obj1.Name)

	//emp_obj2 := Math_Functions.Employee_project{emp_obj1,"Dish-It","skylabs", 12}
	//emp_obj2.Employee_biodata.Basic_info()
	//

	////string literal checks for escape characters and prints the hi in next line.
	//abc:= "hi\nhi"
	//fmt.Println(abc)
	////raw literal (` - backtick) ignores the escape characters and prints the string as is.
	//abc2:=`hi\nhi`
	//fmt.Println(abc2)

	//emp1:= emp_name{"Abishek"}
	//emp1.add_sageit()
	////checking if the value has been changed in real time
	//fmt.Println(emp1.name) //yay!

	//testing interface

	//var1:= Math_Functions.Intnums{1,2}
	//var2:= Math_Functions.Floatnums{1.234, 4.567}
	////Math_Functions.Total_addition(var1)
	//Math_Functions.Total_addition(var2)

	//testing go routines
	//for x:=0;x<10;x++{
	//	go go_routine(x)
	//}
	//var x string
	//fmt.Scanln(&x)
	//fmt.Println(x)

	//scanln reads until first space only but not the entire line of text entered.
	//var x string
	//fmt.Scanln(&x)
	//fmt.Println(x)

	//channels
	//channel1 := make(chan int)
	//go chn_example_func(channel1, 5)
	//go chn_example_func(channel1, 3)

	//outputs from the func chn_example_func are in the channel and we need to retrieve them.
	//a:= <- channel1
	//b:= <- channel1
	// or the below way is fine too
	//a, b:= <-channel1, <-channel1
	//fmt.Println(a,b)

	////calling channel functions declared in Math_Functions package.
	//c:= make(chan string)//declaring a channel
	//go Math_Functions.Pinger(c)
	//go Math_Functions.Ponger(c)
	//go Math_Functions.Printer(c)
	//
	////
	//var input string
	//fmt.Scanln(&input)
	//

	//
	//////testing select in GO in connection with channels.
	//c1:= make(chan string)
	//c2:= make(chan string)
	//go func(){
	//	//for {
	//
	//		time.Sleep(time.Second * 3)
	//		c1 <- "func1"
	//		c1 <- "func1"
	//
	//	//}
	//}()
	//
	//go func(){
	//	//for {
	//	time.Sleep(time.Second * 3)
	//		c2 <- "func2"
	//		c2 <- "func2"
	//
	//	//}
	//}()
	//
	////using select in statement while printing the output from channels.
	//go func(){
	//	for {
	//		select{
	//		case msg1:= <-c1:
	//			fmt.Println(msg1)
	//		case msg2:= <-c2:
	//			fmt.Println(msg2)
	//		//case timevar:= <- time.After(time.Second*2):
	//		//	fmt.Println(timevar)
	//		//default:
	//		//	fmt.Println("nothing is ready")
	//		}
	//	}
	//}()
	//
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println(input)

	//basic http
	//http.HandleFunc("/", home_page)
	//http.HandleFunc("/about", about_page)
	//http.ListenAndServe(":8800",nil)

	//testing rpc server calls example from the book.
	//go server()
	//go client()
	//
	//var input string
	//fmt.Scanln(&input)
	//

	//testing rpc server calls example written by me
	//go server()
	//go client()
	//
	//var input string
	//fmt.Scanln(&input)

	//mutexes example
	//m := new(sync.Mutex)
	//for i := 0; i < 10; i++ {
	//	go sleep(i, m)
	//}
	//var input string
	//fmt.Scanln(&input)

	//m := new(sync.Mutex)
	//
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		m.Lock()
	//		fmt.Println(i, "start")
	//		time.Sleep(time.Second)
	//		fmt.Println(i, "end")
	//		m.Unlock()
	//	}(i)
	//}
	//
	//var input string
	//fmt.Scanln(&input)

	////implementing mutexe on go routines
	//for i:=0; i<10; i++{
	//	//go func(i int){
	//	//	fmt.Println(i)
	//	//}(i) // we should use () when an internal or closure function is not assigned to a variable
	//
	//
	//	somefunc := func() int{
	//		return 10
	//	} // in this case we need not use an calling notation "()" as the closure function is assigned to a variable and the variable is called as an actual function in the next step, "somefunc()".
	//	go fmt.Println(somefunc())
	//}
	//
	//var inp string
	//fmt.Scanln(&inp)

	//const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	//fmt.Println(sample)

	//fmt.Println(SwapRune('b'))

	//x:= "abishek"
	//for _,v:=range x{
	//	fmt.Println(rune(v))
	//}

	//fmt.Println()

	//x := [4]int {1,2,3,4}
	//y := x[0:]
	//y = append(y, 5,6,7,8,9,10)
	//fmt.Println(x)
	//fmt.Println(y)

	//y := make([]int, 4, 7)
	//y = append(y, 5,6,7,8,9,10,11,12,13)
	//fmt.Println(y)
	//fmt.Printf("%c",y)

	//name1 := "Abishek is Abishek"
	//name2 := ""
	//a := []rune(name1)
	//
	//	for j:=len(name1)-1;j>=0;j--{
	//		name2 = name2 + string(a[j])
	//	}
	//fmt.Println(name2)

	//string1 := Math_Functions.ReverOfString("this is some string")
	//fmt.Println(string1)

	//name := Math_Functions.ReverOfString("Abishek")
	//fmt.Println(name)

	//string1:= Math_Functions.Removestringdupes("grammar is grammar")
	//fmt.Println(string1)
	// Math_Functions.ReverOfString("grammar")

	//type wild struct{
	//
	//}
	//
	//type pet struct {
	//
	//}
	//
	//type animal struct{
	//	wild string
	//	pet string
	//}

	//  var x map[int]string

	// x[10]="string"
	// fmt.Println(x)

}

// func SwapRune(r rune) rune {
// 	switch {
// 	case 97 <= r && r <= 122:
// 		return r - 32
// 	case 65 <= r && r <= 90:
// 		return r + 32
// 	default:
// 		return r
// 	}

// }
//
//func main() {
//	fmt.Println(SwapRune('a'))
//}

func sleep(i int, m *sync.Mutex) {
	m.Lock()
	fmt.Println(i, "start")
	time.Sleep(time.Second)
	fmt.Println(i, "end")
	m.Unlock()
}

func go_routine(a int) {
	for i := 0; i < 10; i++ {
		fmt.Println(a, ":", i)
		//amt := time.Duration(rand.Intn(1000))
		//time.Sleep(time.Millisecond * amt)
	}
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//func plustwo() func() int{
//	i:= uint(0)
//	return func() (ret int){
//
//	}
//
//	}
//}

func f2() int {
	return f3()
}

func f3() int {
	return 32
}

//advanced closure function
//func abc() func() int{
//	x:=0
//	return func() (ret int){
//		ret = x
//		x+=1
//		return
//}
//}

func abc1() func() int {
	x := 1
	return func() (ret int) {
		ret = x
		x += 2
		return
	}
}

func recur1(x int) int {
	if x == 0 {
		return 1
	}
	return x * recur1(x-1)
}

//fibonacci
func fib(x int) int {
	if x <= 0 {
		return x
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}

}

func f1(x int) (int, int, bool) {
	a := x + x
	b := x + x + x
	var c bool
	if a != 0 && b != 0 {
		c = true
	} else {
		c = false
	}
	return a, b, c
}

//testing defer
func first() {
	fmt.Println(1)
}

func second() {
	fmt.Println(2)
}

func car_info(a *Car) {
	a.price = math.Round(a.price)
	fmt.Println(a.car_model)
	fmt.Println(a.make)
	fmt.Println(a.year)
	fmt.Println(a.price)
}

//creating a method for type Car

func (a *Car) car_info2() {
	a.price = math.Round(a.price)
	fmt.Println(a.car_model)
	fmt.Println(a.make)
	fmt.Println(a.year)
	fmt.Println(a.price)
}

//creating a method which add "sage-it" for every string
func (a *emp_name) add_sageit() {
	a.name = a.name + "_sage-it"
	fmt.Println(a.name)
}

type emp_name struct {
	name string
}

type car_details struct {
	Car              Car
	Car_vin          string
	Car_reg_year     int
	Car_plate_number string
}

func (a car_details) print_car_details() {
	car_info(&a.Car)
	fmt.Println(a.Car_vin)
	fmt.Println(a.Car_reg_year)
	fmt.Println(a.Car_plate_number)
}

func chn_example_func(c chan int, a int) { //if we want pass the values over a channel we should use chan as an arguement in the function.
	c <- a * 5 //passing a value to the channel. because we are sending a value over the channel we need not return it.
}

//basic http index handle function.
func home_page(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
  			<head>
      			<title>Hello World</title>
  			</head>
  			<body>
      			Hello World!
  			</body>
		</html>`,
	)
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page!")
}

//functions to implement RPC server calls
//defining a struct
type dummytype struct{}

//defining a method using struct type declared above.
func (val *dummytype) anonymous(a string, b *string) {
	*b = a + "_sage_it"
}

func server() {
	rpc.Register(new(dummytype))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result string
	err = c.Call("dummytype.anonymous", string("abc"), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("dummytype.anonymous = ", result)
	}

}

//
//func server() {
//	rpc.Register(new(Server))
//	ln, err := net.Listen("tcp", ":9999")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for {
//		c, err := ln.Accept()
//		if err != nil {
//			continue
//		}
//		go rpc.ServeConn(c)
//	}
//}
//func client() {
//	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	var result int64
//	err = c.Call("Server.Negate", int64(999), &result)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("Server.Negate(999) =", result)
//	}
//}
