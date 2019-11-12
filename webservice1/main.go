package main

import (
	//"fmt"
	"fmt"
	"net/http"
	"webservice1/User_Info"

	//"webservice1/User_Info"
	//"github.com/gorilla/mux"
)

//type Server struct{}


func main(){
	User_Info.Ask_user()
	//a:= &Server{}

	//http.Handle("User_Info", a)
	//http.ListenAndServe(":8000", nil)

	http.HandleFunc("/", home1)
	http.ListenAndServe(":8090", nil)
}

func home1(w http.ResponseWriter, r *http.Request){
	//w.WriteHeader(http.StatusOK)
	//User_Info.Display_user() //this func is unable to display the user info as the func has fmt.printlns which is not a web format. and hence there is nothing displayed on the browser.
	//fmt.Fprintf(w, "hello\n")
	fmt.Fprintln(w,"User_Id: ", User_Info.User_id)
	fmt.Fprintln(w,"Name: ", User_Info.Name)
	fmt.Fprintln(w,"Age: ",User_Info.Age)
	fmt.Fprintln(w,"Technology: ", User_Info.Technology)

}
