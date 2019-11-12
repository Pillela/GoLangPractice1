package main

import (
	"awesomeProject2/All_User_Related_Stuff"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func display_x(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(All_User_Related_Stuff.User_x)
	//w.Write([]byte(`{"message":"get called"}`))


}


func main(){
	fmt.Println("welcome to capital one! Please select one from below options")
	fmt.Println("1 - Enter user info")
	fmt.Println("2 - Exit")
	var choice1 int
	fmt.Scanln(&choice1)
	if choice1 == 1{
		All_User_Related_Stuff.Ask_Info()
	}else {
		return
	}

	fmt.Println("User info succesfully added!")
	All_User_Related_Stuff.Add_additional_user()

	r := mux.NewRouter()

	r.HandleFunc("/User_x", display_x).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))



}
