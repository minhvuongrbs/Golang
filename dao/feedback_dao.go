package dao

import (
	"welcome_robot/models"

	"gopkg.in/mgo.v2/bson"
)

const FeedbackCollection = "Feedback"


func FindAllFeedback() ([]models.Feedback, error) {
	var feedBacks []models.Feedback
	err:=ConnectDatabase().C(FeedbackCollection).Find(bson.M{}).All(&feedBacks)
	return feedBacks, err
}
func InsertFeedback(feedback models.Feedback) (models.Feedback, error) {
	err := ConnectDatabase().C(FeedbackCollection).Insert(&feedback)
	return feedback, err
}


