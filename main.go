package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovojoe/Doc-Up/models"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to DocUp. Easily access all your documents!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter := mux.NewRouter().StrictSlash(true)
	/// http.HandleFunc("/", homePage)
	// add our superheroes route and map it to our
	// returnAllsuperheroes function like so
	/// http.HandleFunc("/superheroes", returnAllSuperheroes)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/returnsingleuser/{id}", models.ReturnSingleUser)
	myRouter.HandleFunc("/returnsingledocument/{id}", models.ReturnSingleDocument)
	myRouter.HandleFunc("/returndocuments", models.ReturnAllDocuments)
	 // NOTE: Ordering is important here! This POST method has to be defined before
    // the other `/user` endpoint. 
    myRouter.HandleFunc("/createuser", models.CreateNewUser).Methods("POST")
    myRouter.HandleFunc("/createdocument", models.CreateNewDocument).Methods("POST")
    myRouter.HandleFunc("/updatedocument/{id}", models.UpdateDocument).Methods("PUT")
    myRouter.HandleFunc("/deleteuser/{id}", models.DeleteUser).Methods("DELETE")
	

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main (){
	//Adding dummy user and document data
	models.Documents = []models.Document{
		{Id: "1", Name: "Driver's license", DateCreated: "25th August"},
	}
	models.Users=[]models.User{
		{Id: "1",Name: "Kratos", Email: "kratos@zeus.die",Password: "Boooooy"},
	}
handleRequests()
}


