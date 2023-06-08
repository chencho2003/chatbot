package controller

import (
	"encoding/json"
	"myProject/model"
	"myProject/utils/httpResp"
	"net/http"
	"myProject/datastore/postgres"
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

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	//storing the dataStored
	var stud model.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	//checking user credaintials
	loginErr := stud.CheckIn()

	if loginErr != nil {
		// switch loginErr {
		// case sql.ErrNoRows:
		// 	httpResp.RespondWithError(w, http.StatusNotFound, "user not found")
		// 	return
		// default:
		// 	httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		// 	return

		// }
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return 

	}
	//statusok 200
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Success"})
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
