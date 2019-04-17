package route

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"welcomerobot-api/dao"
	. "welcomerobot-api/models"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	user, err := dao.GetUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}
func InsertUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	var userInfo UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	_, user, err := dao.InsertUser(userInfo)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, user)
}
func CheckDuplicationName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var name string
	var bodyContent = make(map[string]string)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(body, &bodyContent)
	name = bodyContent["name"]
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = dao.GetUserByName(name)
	m := make(map[string]bool)
	if err != nil {
		//not found name in database
		m["is_duplicate"] = false
		respondWithJson(w, http.StatusOK, m)
		return
	}
	m["is_duplicate"] = true
	respondWithJson(w, http.StatusOK, m)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.DeleteUser(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
