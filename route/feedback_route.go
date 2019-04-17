package route

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	"welcomerobot-api/dao"
	. "welcomerobot-api/models"
)

func InsertFeedback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var feedback Feedback
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	feedback.FeedbackID = bson.NewObjectId()
	feedback.CreatedAT = time.Now()
	fb, err := dao.InsertFeedback(feedback)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, fb)
}
func GetAllFeedbacks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	feedBacks, err := dao.GetAllFeedbacks()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, feedBacks)
}
