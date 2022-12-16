package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var Users []User
var Db *gorm.DB

//Creating a new user

func (h handler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	unmarshalled := json.Unmarshal(reqBody, &user)

	// update the global Users array to include
	// the new Superhero
	fmt.Println(unmarshalled)

	//Users = append(Users, user)
	if result := Db.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(user)
}

// Get user docs
func (h handler) GetUsersDocs(w http.ResponseWriter, r *http.Request) {
	var users []User
	//err := db.Model(&models.User{}).Preload("Documents").Find(&users).Error
	if results := Db.Preload("Documents").Find(&users); results.Error != nil {
		fmt.Println(results.Error)
	}

	fmt.Println("got users")
	json.NewEncoder(w).Encode(users)

}

// Fetch a single user and all their documents
func (h handler) ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var singleUser User

	//Get the first user whose ID matches the provided ID
	Db.First(&singleUser, key)
	if result := Db.First(&singleUser, key); result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(w).Encode(singleUser)

}

// Deleting a user by ID
func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the user we
	// wish to delete
	id := vars["id"]
	var deletedUser User

	if result := Db.First(&deletedUser, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	Db.Delete(&deletedUser)
	json.NewEncoder(w).Encode(deletedUser)


}

// Creating a user model/struct containing all the attributes of  a user
type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Documents []Document
}

type Document struct {
	gorm.Model
	Name   string
	UserID uint
}
