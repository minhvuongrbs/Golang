package dao

import (
	"log"
	"time"
	"welcome_robot/models"

	"gopkg.in/mgo.v2/bson"
)

const FeedbackCollection = "Feedback"


func FindAllFeedback() ([]models.Feedback, error) {
	var fbs []models.Feedback
	c := ConnectDatabase().C(FeedbackCollection)
	err := c.Find(nil).All(&fbs)
	if err != nil {
		log.Fatal(err.Error())
	}
	return fbs, err
}
func InsertFeedback(fb models.Feedback) (models.Feedback, error) {
	feedback := models.Feedback{
		FeedbackID: bson.NewObjectId(),
		Comment:    fb.Comment,
		Rating:     fb.Rating,
		CreatedAT:  time.Now(),
	}
	err := ConnectDatabase().C(FeedbackCollection).Insert(feedback)
	if err != nil {
		log.Fatal("Insert Fail")
	}
	return feedback, err
}


