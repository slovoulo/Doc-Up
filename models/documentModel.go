package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var Documents []Document

//Creating a new document
func createNewDocument(w http.ResponseWriter, r *http.Request) {
    // get the body of the  POST request
    // unmarshal this into a new Document struct
    // append this to the Documents array.     
    reqBody, _ := ioutil.ReadAll(r.Body)
	var document Document 
    unmarshalled:=json.Unmarshal(reqBody, &document)

	// update the global Users array to include
    // the new Superhero
	fmt.Println(unmarshalled)
	
    Documents = append(Documents, document)
    json.NewEncoder(w).Encode(document)
}

//Fetch a single document
func returnSingleDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all documents
	// if the document.Id equals the key we pass in
	// return the document encoded as JSON
	for _, document := range Documents {
		if document.Id == key {
			json.NewEncoder(w).Encode(document)
		}
	}
}

//Fetch all documents
func returnAllDocuments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDocuments")

	//The call to json.NewEncoder(w).Encode(article) does the job of encoding our documents array into a JSON string and then writing as part of our response.
	json.NewEncoder(w).Encode(Documents)
}

//Update document
func updateDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEvent Document
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, document := range Documents {
		if document.Id == id {
	
			document.Name = updatedEvent.Name
			document.Id = updatedEvent.Id
			document.DateCreated = updatedEvent.DateCreated
			
		
			Documents[i] = document
			json.NewEncoder(w).Encode(document)
		}
	}
	
	}

//Deleting a document
func deleteDocument(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the doc we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all Documents
    for index, document := range Documents {
        // if our id path parameter matches one of our
        // articles
        if document.Id == id {
            // updates  Users array to remove the 
            // user
            Documents = append(Documents[:index], Documents[index+1:]...)
        }
    }

}

// Creating a user model/struct containing all the attributes of  a user
type Document struct {
	
	Id           string `json: "id"`
	Name         string	`json:"name"`
	DateCreated         string	`json:"dateCreated"`
	
	
}