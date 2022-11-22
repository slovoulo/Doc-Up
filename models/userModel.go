package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var Users []User

//Creating a new user
func createNewUser(w http.ResponseWriter, r *http.Request) {
    // get the body of the  POST request
    // unmarshal this into a new User struct
    // append this to the Users array.     
    reqBody, _ := ioutil.ReadAll(r.Body)
	var user User 
    unmarshalled:=json.Unmarshal(reqBody, &user)

	// update the global Users array to include
    // the new Superhero
	fmt.Println(unmarshalled)
	
    Users = append(Users, user)
    json.NewEncoder(w).Encode(user)
}

//Fetch a single user
func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all users
	// if the user.Id equals the key we pass in
	// return the user encoded as JSON
	for _, user := range Users {
		if user.Id == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}

//Deleting a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the user we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all users
    for index, user := range Users {
        // if our id path parameter matches one of our
        // articles
        if user.Id == id {
            // updates  Users array to remove the 
            // user
            Users = append(Users[:index], Users[index+1:]...)
        }
    }

}

// Creating a user model/struct containing all the attributes of  a user
type User struct {
	
	Id           string `json: "id"`
	Name         string	`json:"name"`
	Email         string	`json:"email"`
	Password  string	`json:"string"`
	
}