package main

import (
	"github.com/gorilla/mux"
	app "goProject/mysql-api/app"
	"goProject/mysql-api/db"
	"log"
	"net/http"
)

//func setupRouter(router *mux.Router){
//	router.
//		Methods("POST").
//		Path("/endpoint").
//		HandlerFunc(postFunction)
//}
//
//func postFunction(w http.ResponseWriter, r *http.Request){
//	database, err := db.CreateDatabase()
//	if err != nil {
//		log.Fatal("Database connection failed")
//	}
//
//	_, err = database.Exec("INSERT INTO `test` (name) VALUES ('myname')INSERT INTO `test` (name) VALUES ('myname')")
//	if err != nil {
//		log.Fatal("Database INSERT failed")
//	}
//
//	log.Println("you called a thing")
//}

func main(){
	database, err := db.CreateDatabase()

	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router: mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	//router := mux.NewRouter().StrictSlash(true)
	//
	//setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080",app.Router))
}