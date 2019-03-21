package route

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"welcome_robot/dao"
	. "welcome_robot/models"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	user, err := dao.FindUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}
func InsertUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	var userInfor UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfor); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := dao.InsertUser(userInfor)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.RemoveUser(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
