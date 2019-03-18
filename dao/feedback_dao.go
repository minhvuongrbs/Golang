package dao

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
	"welcome_robot/models"
)

const (
	FeedbackCollection = "Feedback"
)

func InsertFeedback(fb models.Feedback) (models.Feedback, error) {
	feedback := models.Feedback{
		FeedbackID:bson.NewObjectId(),
		Comment:fb.Comment,
		Rating:fb.Rating,
		CreatedAT:time.Now(),
	}
	err := ConnectDatabase().DB(DatabaseName).C(FeedbackCollection).Insert(feedback)
	if err != nil {
		log.Fatal("Insert Fail")
	}
	return feedback, err
}

func GetAllFeedback() ([]models.Feedback, error) {
	var fbs []models.Feedback
	c := ConnectDatabase().DB(DatabaseName).C(FeedbackCollection)
	err := c.Find(nil).All(&fbs)
	if err != nil {
		log.Fatal(err.Error())
	}
	return fbs,err
}