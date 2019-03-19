package route

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	"welcome_robot/dao"
	. "welcome_robot/models"
)

func InsertVideoTime(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var videoTime VideoTime
	if err := json.NewDecoder(r.Body).Decode(&videoTime); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	videoTime.VideoTimeID = bson.NewObjectId()
	videoTime.IsPause = false
	videoTime.TimeStamp = 0
	videoTime.TimeStart = time.Now()
	if err := dao.InsertVideoTime(videoTime); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videoTime)
}
func GetVideoTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoTime, err := dao.FindVideoTimeById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid video time ID")
		return
	}
	respondWithJson(w, http.StatusOK, videoTime)
}
func UpdateVideoTime(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //execute when updateVideoTime function finished (defer)
	var videoTime VideoTime
	if err := json.NewDecoder(r.Body).Decode(&videoTime); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateVideoTime(videoTime); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func GetAllVideoTime(w http.ResponseWriter, r *http.Request) {
	videoTimes, err := dao.FindAllVideoTime()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videoTimes)
}
func DeleteVideoTime(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.DeleteVideoTime(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}