package route

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	"welcomerobot-api/dao"
	. "welcomerobot-api/models"
)

type DetailSession struct {
	SessionId   bson.ObjectId `bson:"session_id" json:"session_id"`
	Supporter   User          `bson:"supporter" json:"supporter"`
	User        User          `bson:"user" json:"user"`
	CheckInTime time.Time     `bson:"check_in_time" json:"check_in_time"`
}

func InsertSession(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var session Session
	var user User
	var userInfo UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if userInfo.Permission != 2 {
		session.CheckInTime = time.Now()
	}
	isInsert, user, err := dao.InsertUser(userInfo)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	session.SupporterID = getSupporterId()
	session.UserID = user.UserID
	if isInsert {
		session.SessionID = bson.NewObjectId()
		if err := dao.InsertSession(session); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		session2, err := dao.GetSessionByUserID(user.UserID.Hex())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		session.SessionID = session2.SessionID
		if err := dao.UpdateSession(session); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJson(w, http.StatusOK, session)
}
func getSupporterId() bson.ObjectId {
	//TODO: replace by query method
	//create default user for testing
	return bson.ObjectIdHex("5c98538a31ce9717b85de4aa")
}
func GetAllSessions(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var sessions [] Session
	var supporter User
	var user User
	sessions, err := dao.GetAllSessions()
	detailSessions := make([]DetailSession, len(sessions))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	for i := 0; i < len(sessions); i++ {
		supporter, err = dao.GetUserById(sessions[i].SupporterID.Hex())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, errors.New("server not found supporter id").Error())
			return
		}
		user, err = dao.GetUserById(sessions[i].UserID.Hex())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, errors.New("server not found user id").Error())
			return
		}
		detailSessions[i].SessionId = sessions[i].SessionID
		detailSessions[i].Supporter = supporter
		detailSessions[i].User = user
		detailSessions[i].CheckInTime = sessions[i].CheckInTime
	}
	respondWithJson(w, http.StatusOK, detailSessions)
}
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var session Session
	session, err := dao.GetSessionById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = dao.DeleteUser(session.UserID.Hex())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = dao.DeleteVideoTimeBySessionId(session.SessionID.Hex())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := dao.DeleteSession(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func GetSessionByUserId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	respondWithError(w, http.StatusBadRequest, errors.New("not found required").Error())
}
