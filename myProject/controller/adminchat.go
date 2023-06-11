package controller

import (
	"encoding/json"
	"myProject/model"
	"myProject/utils/httpResp"
	"net/http"
	"database/sql"
	"myProject/datastore/postgres"
)
func VerifyCookie(w http.ResponseWriter,r *http.Request) bool {
	cookie, err := r.Cookie("admin-cookie")
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			httpResp.RespondWithError(w,http.StatusSeeOther,"cookie not set")
		default:
			httpResp.RespondWithError(w,http.StatusInternalServerError,"internal server error")
		}
		return false
	}


	if cookie.Value != "#@chatbotcoolazha" {
		httpResp.RespondWithError(w,http.StatusSeeOther,"invalid cookie")
		return false
	}
	return true
}

func TeachingBot(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w,r) {
		return 
	}
	var QNA model.Bot
	// fmt.Println(stud)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&QNA)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
	}
	// fmt.Println(stud)

	saveErr := QNA.Put()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
	} else {
		//status crested 201
		httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Question and Answer added"})
	}
	// fmt.Fprintf(w, "add student handler")
}
func Deleting(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w,r) {
		return 
	}
	var QNA model.Bot
	// fmt.Println(stud)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&QNA)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
	}
	// fmt.Println(stud)

	saveErr := QNA.DeleteData()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
	} else {
		//status crested 201
		httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "data deleted"})
	}
	// fmt.Fprintf(w, "add student handler")
}

func Chat(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w,r) {
		return 
	}
	var QNA model.Bot
	// fmt.Println(stud)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&QNA)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
	}
	getErr := QNA.Accessing()

	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Sorry.. I could not find the answer")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, QNA)
	}
}
func AllData(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w,r) {
		return 
	}
	//cookie verification


	datas, getErr := model.GetAllData()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, datas)

}
func UpdateQNA(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w,r) {
		return 
	}
	var QNA model.Bot
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&QNA)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	// Prepare the SQL statement
	stmt, err := postgres.Db.Prepare("UPDATE bot SET question = $1 , answer = $2  WHERE id = $3")
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Execute the prepared statement with the parameter values
	_, err = stmt.Exec(QNA.Question, QNA.Answer, QNA.Id)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the updated student details
	httpResp.RespondWithJSON(w, http.StatusOK, QNA)
}