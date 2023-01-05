package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/slovojoe/Doc-Up/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	myRouter.HandleFunc("/returnsingledocument/{id}", models.ReturnSingleDocument).Methods("GET")
	myRouter.HandleFunc("/returndocuments", models.ReturnAllDocuments).Methods("GET")
	// NOTE: Ordering is important here! This POST method has to be defined before
	// the other `/user` endpoint.
	// myRouter.HandleFunc("/createuser", h.CreateNewUser).Methods("POST")
	

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// Adding dummy user and document data
var (
	documents = []models.Document{
		{Name: "Driver's license", UserID: 1},
		{Name: "passport", UserID: 1},
	}
	users = []models.User{
		{Name: "Kratos", Email: "kratos@zeus.die", Password: "Boooooy"},
	}
)

func GetAll(db *gorm.DB) []*models.User {
	var users []*models.User
	//err := db.Model(&models.User{}).Preload("Documents").Find(&users).Error
	//results:=db.Preload("Documents").Find(&users)
	results := db.Find(&users)
	fmt.Println(results)
	fmt.Println("got users")
	return users
}

func main() {
	
	var err error
	myRouter := mux.NewRouter().StrictSlash(true)

	//Loading environment variables
	//dialect:= os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//Database connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, dbPort, user, dbName, password)

	//Opening connection to DB
	models.Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	h := models.New(models.Db)

	myRouter.HandleFunc("/createuser", h.CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/getall", h.GetUsersDocs).Methods("GET")
	myRouter.HandleFunc("/createdocument", h.CreateNewDocument).Methods("POST")
	myRouter.HandleFunc("/returnsingleuser/{id}", h.ReturnSingleUser).Methods("GET")
	myRouter.HandleFunc("/deleteuser/{id}", h.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/updatedocument/{id}", h.UpdateDocument).Methods("PUT")
	myRouter.HandleFunc("/deletedocument/{id}", h.DeleteDocument).Methods("DELETE")

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Successfully connected to a database")
	}

	//Close connrction to db when main function finishes
	//*defer db.Close()

	models.Db.AutoMigrate(&models.Users, &models.Documents)
	//models.Db.Migrator().DropTable(&models.User{})
	models.Db.Migrator().CreateTable(&models.User{})
	//models.Db.Migrator().DropTable(&models.Document{})

	//Make migrations to the DB if they have not been made

	models.Db.Migrator().CreateTable(&models.Document{})

	// 	models.Db.Create(&models.User{Name: "Ellie",Email: "Ellie@lous",
	// 	Documents: []models.Document{{Name:"Ellies ID",},{Name: "Ellies Guitar "}},
	// })

	// models.Db.Debug().Save(&models.User{
	// 	Name:"Beans",
	// 	Documents: []models.Document{
	// 		{Name: "ID"},
	// 		{Name: "DL"},
	// 	},
	// })

	//get users
	//GetAll(models.Db)

	// for i := range documents{models.Db.Create(documents[i])}
	// for i := range users{models.Db.Create(&users[i])}
	//handleRequests()
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
