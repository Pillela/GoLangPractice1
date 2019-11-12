package User_Info

import "fmt"

var Name, Technology string
var User_id, Age int

func Ask_user(){
	fmt.Println("Hello User!")
	fmt.Println("Whats your user_id ?")
	fmt.Scanln(&User_id)
	fmt.Println("Whats your name ?")
	fmt.Scanln(&Name)
	fmt.Println("Whats your Technology ?")
	fmt.Scanln(&Technology)
	fmt.Println("Whats your age ?")

	fmt.Scanln(&Age)

}

