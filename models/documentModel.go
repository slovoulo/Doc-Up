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
func (h handler) CreateNewDocument(w http.ResponseWriter, r *http.Request) {
    // get the body of the  POST request
    // unmarshal this into a new Document struct
    // append this to the Documents array.     
    reqBody, _ := ioutil.ReadAll(r.Body)
	var document Document 
    unmarshalled:=json.Unmarshal(reqBody, &document)


	fmt.Println(unmarshalled)
	
    if result := Db.Create(&document); result.Error != nil {
		fmt.Println(result.Error)
	}
    json.NewEncoder(w).Encode(document)
}

//Fetch a single document
func ReturnSingleDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	var singleDoc  []Document

	//Find the first value that matches the given parameters (ID)
	Db.First(&singleDoc, key)
	json.NewEncoder(w).Encode(&singleDoc)
}


//Fetch all documents
func ReturnAllDocuments(w http.ResponseWriter, r *http.Request) {
	var allDocs []Document
	Db.Find(&allDocs)
	json.NewEncoder(w).Encode(&allDocs)
}

//Update document by id
func (h handler)UpdateDocument(w http.ResponseWriter, r *http.Request) {

	// once again, we will need to parse the path parameters
	var updatedDoc Document
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedDoc)
	var document Document
	vars := mux.Vars(r)
	id := vars["id"]

	
	
	if result := Db.First(&updatedDoc, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	document.Name=updatedDoc.Name

	Db.Save(&document)
	json.NewEncoder(w).Encode(&updatedDoc)
}
	
	//}

//Deleting a document
// func DeleteDocument(w http.ResponseWriter, r *http.Request) {
//     // once again, we will need to parse the path parameters
//     vars := mux.Vars(r)
//     // we will need to extract the `id` of the doc we
//     // wish to delete
//     id := vars["id"]

//     // we then need to loop through all Documents
//     for index, document := range Documents {
//         // if our id path parameter matches one of our
//         // articles
//         if document.Id == id {
//             // updates  Users array to remove the 
//             // user
//             Documents = append(Documents[:index], Documents[index+1:]...)
//         }
//     }

// }

// Creating a user model/struct containing all the attributes of  a user
