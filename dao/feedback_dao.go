package dao

import (
	"welcomerobot-api/models"

	"gopkg.in/mgo.v2/bson"
)

func GetAllFeedbacks() ([]models.Feedback, error) {
	var feedBacks []models.Feedback
	err := ConnectDatabase().C(FeedbackCollection).Find(bson.M{}).All(&feedBacks)
	return feedBacks, err
}
func InsertFeedback(feedback models.Feedback) (models.Feedback, error) {
	err := ConnectDatabase().C(FeedbackCollection).Insert(&feedback)
	return feedback, err
}
