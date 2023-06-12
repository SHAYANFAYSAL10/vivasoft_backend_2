package user_api

import (
	"VIVASOFT2/src/config"
	"VIVASOFT2/src/entities"
	"VIVASOFT2/src/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		users, err2 := userModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, users)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var insert_user entities.Insert_user
	err := json.NewDecoder(request.Body).Decode(&insert_user)
	if err != nil {
		fmt.Println(err)
	}
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		err2 := userModel.Create(&insert_user)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, insert_user)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var insert_user entities.Insert_user
	err := json.NewDecoder(request.Body).Decode(&insert_user)
	if err != nil {
		fmt.Println(err)
	}
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		_, err2 := userModel.Update(&insert_user)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, insert_user)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		RowsAffected, err2 := userModel.Delete_user(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]int64{
				"Rows Affected User": RowsAffected,
			})
		}
		RowsAffected2, err2 := userModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]int64{
				"Rows Affected Activity_logs": RowsAffected2,
			})
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
