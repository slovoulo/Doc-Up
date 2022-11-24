package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/slovojoe/Doc-Up/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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

	//Loading environment variables
	//dialect:= os.Getenv("DIALECT")
	host:= os.Getenv("HOST")
	dbPort:= os.Getenv("DBPORT")
	user:= os.Getenv("USER")
	dbName:= os.Getenv("NAME")
	password:= os.Getenv("PASSWORD")

	//Database connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, dbPort, user, dbName, password)

	//Opening connection to DB
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err !=nil{
		log.Fatal(err)

	}else{fmt.Println("Successfully connected to a database")}

	//Close connrction to db when main function finishes
	//*defer db.Close()

	//Make migrations to the DB if they have not been made
	db.AutoMigrate(&models.Document{})
	db.AutoMigrate(&models.User{})

	// //Adding dummy user and document data
	// models.Documents = []models.Document{
	// 	{Id: "1", Name: "Driver's license", DateCreated: "25th August"},
	// }
	// models.Users=[]models.User{
	// 	{Id: "1",Name: "Kratos", Email: "kratos@zeus.die",Password: "Boooooy"},
	// }
handleRequests()
}


