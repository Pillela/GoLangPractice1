package All_User_Related_Stuff

import ("fmt")
//
//var SSN int
//var Name, Address string
//var Salary float64

type User struct{
	Name string
	Address string
	Salary float64
	SSN int
}

//var User_x []User
var User_x []User
func Ask_Info(){
	fmt.Println("Please provide your Address in quotes")
	var address string
	fmt.Scanf("%q",&address)

	fmt.Println("Please provide your SSN")
	var ssn int
	fmt.Scan(&ssn)

	fmt.Println("Please provide your Salary")
	var salary float64
	fmt.Scanln(&salary)

	fmt.Println("Please provide your Name in quotes")
	var name string
	fmt.Scanf("%q",&name)

	User_x = append(User_x, User{Name: name, Salary:salary, Address:address, SSN:ssn})

}

//func adding_more_users(){
//	User_x = User_x.append(Ask_Info())
//}

var input2 int

func Add_additional_user() {
	fmt.Println("Would you like to add another user info ? please enter 1 or 0.")
	fmt.Scan(&input2)
	if input2==1{
		fmt.Println("checked2")
		Ask_Info()
	}else{
		return
	}
}

