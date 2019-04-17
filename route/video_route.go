package route

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"welcomerobot-api/dao"
	. "welcomerobot-api/models"
)

func InsertVideo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var video Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	video.VideoID = bson.NewObjectId()
	if err := dao.InsertVideo(video); err != nil {
		respondWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, video)

}
func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	videos, err := dao.GetAllVideos()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videos)
}
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.DeleteVideo(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
