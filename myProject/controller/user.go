package controller

import (
	"database/sql"
	"encoding/json"
	"myProject/model"
	"myProject/utils/httpResp"
	"net/http"
	"myProject/datastore/postgres"
	"fmt"
	"time"
)


func Adduser(w http.ResponseWriter, r *http.Request) {
	var stud model.User
	// fmt.Println(stud)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
	}
	// fmt.Println(stud)

	saveErr := stud.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
	} else {
		//status crested 201
		httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User data added"})
	}
	// fmt.Fprintf(w, "add student handler")
}

// use capital function Name to be able call outside
var admin model.User

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	const (
		StatusMyCustomCode = 480
	)
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil{
		httpResp.RespondWithError(w,http.StatusBadRequest,"invalid json body")
		return 
	}
	defer r.Body.Close()
	email := admin.Email
	var admin2 model.User
	loginErr := admin2.Check(email)

	if loginErr != nil{
		switch loginErr{
		case sql.ErrNoRows:
			httpResp.RespondWithError(w,http.StatusUnauthorized,"invalid login")
		default:
			httpResp.RespondWithError(w,http.StatusBadRequest,"error in database")
		}
		return 
	}
	fmt.Println(admin.Password,"requst")
	fmt.Println(admin2.Password,"database")
	if admin.Password != admin2.Password{
		httpResp.RespondWithError(w,http.StatusUnauthorized,"invalid login")
		return 
	}
	cookie := http.Cookie{
		Name: "admin-cookie",
		Value: "#@chatbotcoolazha",
		Expires: time.Now().Add(30 * time.Minute),
		Secure: true,
	}
	//set cookie and send back to client
	http.SetCookie(w,&cookie)
	if admin.Email == "admin@gmail.com"{
		httpResp.RespondWithError(w,StatusMyCustomCode,"admin")
		return 
	}
	//create a cookie

	httpResp.RespondWithJSON(w,http.StatusOK,map[string]string{"message":"successful"})
}
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var stud model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	// Prepare the SQL statement
	stmt, err := postgres.Db.Prepare("UPDATE userdata SET password = $1 WHERE email = $2")
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Execute the prepared statement with the parameter values
	_, err = stmt.Exec(stud.Password, stud.Email)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the updated student details
	httpResp.RespondWithJSON(w, http.StatusOK, stud)
}
