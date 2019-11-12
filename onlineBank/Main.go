package main

import (
	"fmt"
	"runtime"
)

type Account_Info struct {
	Account_id int
	User_Name string
	User_pwd string
}

func main(){
	fmt.Println("Welcome to ABC Bank!")
	fmt.Println("Please select an option from below:")
	fmt.Println("1- Create a new account")
	fmt.Println("2- Login using an existing account")
	fmt.Println("3- Exit the application")
	var User_choice int
	fmt.Scanln(&User_choice)

	if User_choice=1{
		Account_Creation()
	}else if User_choice=2{
		Account_Login()
	}else if User_choice=3{

	}else{
		fmt.Println("Sorry, incorrect choice. Please select one of the above options.")
	}






}
