package main

import (
	"database/sql"
	"log"
)

type Customer struct {

	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name`
	Address string `json:"address"`
	City string `json:"city"`
	Account_Balance float64 `json:account_balance`
}




func main() {
}