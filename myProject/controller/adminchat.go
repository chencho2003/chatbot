package controller

import (
	"encoding/json"
	"myProject/model"
	"myProject/utils/httpResp"
	"net/http"
	"database/sql"
)

func TeachingBot(w http.ResponseWriter, r *http.Request) {
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
// func Chat(w http.ResponseWriter, r *http.Request) {
// 	var QNA model.Bot
// 	// fmt.Println(stud)
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&QNA)
// 	if err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
// 	}
// 	// fmt.Println(stud)

// 	saveErr := QNA.Accessing()
// 	if saveErr != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
// 	} else {
// 		//status crested 201
// 		httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "data deleted"})
// 	}
// 	// fmt.Fprintf(w, "add student handler")
// }

func Chat(w http.ResponseWriter, r *http.Request) {
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
			httpResp.RespondWithError(w, http.StatusNotFound, "could not find the answer")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, QNA)
	}
}


