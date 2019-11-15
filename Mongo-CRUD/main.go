package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//creating mongo session
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}

//UserController export
type UserController struct {
	session *mgo.Session
}

//NewUserController session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//Car struct
type Car struct {
	Make  string        `json:"make" bson:"Make"`
	Model string        `json:"model" bson:"Model"`
	Price float32       `json:"price" bson:"Price"`
	Id    bson.ObjectId `json:"id" bson:"_id"`
}

// ViewAllCars export
func (uc UserController) ViewAllCars(w http.ResponseWriter, r *http.Request) {
	u := []Car{}
	if err := uc.session.DB("Database1").C("Cars").Find(nil).All(&u); err != nil {
		fmt.Println("Error1")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Connected to ")
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

//CreateCar export
func (uc UserController) CreateCar(w http.ResponseWriter, r *http.Request) {
	u := Car{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	fmt.Println(u)
	uc.session.DB("Database1").C("Cars").Insert(u)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// ViewCarById export
func (uc UserController) ViewCarById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("_id")
	// fmt.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// fmt.Println(id)
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	u := Car{}
	if err := uc.session.DB("Database1").C("Cars").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteCar export
func (uc UserController) DeleteCar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("_id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("Database1").C("Cars").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deleted car", oid, "\n")
}

//UpdateCar export
func (uc UserController) UpdateCar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("_id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//object id
	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)

	u := Car{}
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	uid := bson.M{"_id": oid}
	fmt.Println(uid)

	uc.session.DB("Database1").C("Cars").Update(uid, bson.M{"Make": u.Make, "Model": u.Model, "Price": u.Price})
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func main() {
	uc := NewUserController(getSession())
	http.HandleFunc("/viewall", uc.ViewAllCars)
	http.HandleFunc("/view", uc.ViewCarById)
	http.HandleFunc("/create", uc.CreateCar)
	http.HandleFunc("/delete", uc.DeleteCar)
	http.HandleFunc("/update", uc.UpdateCar)
	log.Fatal(http.ListenAndServe(":8098", nil))
}
