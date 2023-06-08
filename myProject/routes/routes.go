package routes

import (
	"fmt"
	"myProject/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitializerRoutes() {
	//creating a new router
	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Adduser).Methods("POST")
	//creating routes using handler for login...
	router.HandleFunc("/login", controller.Loginhandler).Methods("POST")
	//updating password
	router.HandleFunc("/update", controller.UpdatePassword).Methods("PUT")
	//teaching bot
	router.HandleFunc("/teach", controller.TeachingBot).Methods("POST")
	//deleting data
	router.HandleFunc("/delete", controller.Deleting).Methods("DELETE")
	//accessing data
	router.HandleFunc("/chat", controller.Chat).Methods("POST")


	//serving the static files
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	// start the http server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// log.Println("Application running on port 8080")

	//OR
	//log.Fatal(http.ListenAndServe(":8080", router))
}
