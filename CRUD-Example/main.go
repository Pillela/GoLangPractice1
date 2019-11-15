package main

import (
	"database/sql"
	//"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Car struct {
	ID    int
	Make  string
	Model string
	Cost  float32
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root1234*"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// //Show func
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nid := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM cars WHERE id=?", nid)
	if err != nil {
		panic(err.Error())
	}
	car := Car{}
	for selDB.Next() {
		var id int
		var make, model string
		var cost float32
		err = selDB.Scan(&id, &make, &model, &cost)
		if err != nil {
			panic(err.Error())
		}
		car.ID = id
		car.Make = make
		car.Model = model
		car.Cost = cost
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"ID": }`))
	// fmt.Fprintf("Hello")

	// tmpl.ExecuteTemplate(w, "Show", car)
	defer db.Close()
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM cars ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	car := Car{}
	result := []Car{}
	for selDB.Next() {
		var id int
		var make, model string
		var cost float32
		err = selDB.Scan(&id, &make, &model, &cost)
		if err != nil {
			panic(err.Error())
		}
		car.ID = id
		car.Make = make
		car.Model = model
		car.Cost = cost
		result = append(result, car)
	}

	// tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id": "world"}`))
}

func main() {
	// log.Println("Server started on: http://localhost:8080")
	// s := &server{}
	// http.HandleFunc("/", Index)
	http.HandleFunc("/", Index)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
