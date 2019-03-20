package route

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	"welcome_robot/dao"
	. "welcome_robot/models"
)

func InsertFeedback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var feedback Feedback
	feedback.FeedbackID = bson.NewObjectId()
	feedback.CreatedAT = time.Now()
	fb, err := dao.InsertFeedback(feedback)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, fb)
}
func GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	feedBacks, err := dao.FindAllFeedback()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, feedBacks)
}
