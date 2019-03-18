package _go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
	. "welcome_robot/config"
	"welcome_robot/dao"
	. "welcome_robot/models"
)

var config = Config{}
var videoDAO = dao.VideoDAO{}
var videoTimeDAO = dao.VideoTimeDAO{}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).Handler(handler)

	}

	return router
}
func init() {
	config.Read()
	videoDAO.Server = config.Server
	videoDAO.Database = config.Database
	videoDAO.Connect()
	videoTimeDAO.Server = config.Server
	videoTimeDAO.Database = config.Database
	videoTimeDAO.Connect()
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/wr/v1/home",
		Index,
	},
	Route{
		"Create User",
		"POST",
		"/wr/v1/users",
		InsertUser,
	},
	Route{
		"Remove User",
		"DELETE",
		"/wr/v1/users/{id}",
		RemoveUser,
	},
	Route{
		"GetAllSession",
		"GET",
		"/wr/v1/sessions",
		GetAllSession,
	},
	Route{
		"GetSession",
		"GET",
		"/wr/v1/sessions/{id}",
		GetSession,
	}, Route{
		"InsertSession",
		"POST",
		"/wr/v1/sessions",
		InsertSessions,
	},
	Route{
		"RemoveSession",
		"DELETE",
		"/wr/v1/sessions/{id}",
		RemoveSessions,
	},
	Route{
		"GetAllVideo",
		"GET",
		"/wr/v1/videos",
		GetAllVideo,
	},
	Route{
		"Insert Video",
		"POST",
		"/wr/v1/videos",
		InsertVideo,
	},
	Route{
		"Get Video Time",
		"GET",
		"/wr/v1/videotimes/{id}",
		GetVideoTime,
	}, Route{
		"Delte Video",
		"DELETE",
		"/wr/v1/videos/{id}",
		DeleteVideo,
	},
	Route{
		"create video time",
		"POST",
		"/wr/v1/videotimes",
		InsertVideoTime,
	}, Route{
		"Get all video time",
		"GET",
		"/wr/v1/videotimes",
		GetAllVideoTime,
	},
	Route{
		"update video time",
		"PUT",
		"/wr/v1/videotimes",
		UpdateVideoTime,
	},
	Route{
		"delete video time",
		"DELETE",
		"/wr/v1/videotimes/{id}",
		DeleteVideoTime,
	},
	Route{
		"Insert feedback",
		"POST",
		"/wr/v1/feedbacks",
		InsertFeedback,
	},Route{
		"Get All Feedback",
		"GET",
		"/wr/v1/feedbacks",
		GetAllFeedback,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	var userInfor UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfor); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user, err := dao.InsertUser(userInfor)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.RemoveUser(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func InsertSessions(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var session Session
	var user User
	var userInfor UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfor); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	session.SessionID = bson.NewObjectId()
	if userInfor.Permission != 2 {
		session.CheckInTime = time.Now()
		//respondWithJson(w, http.StatusCreated, user2)
	}
	user, err := dao.InsertUser(userInfor)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//respondWithJson(w, http.StatusCreated, user)
	session.SupporterID = getSupporterId()
	session.UserID = user.UserID
	if err := dao.InsertSession(session); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, session)
}

func getSupporterId() bson.ObjectId {
	return bson.ObjectIdHex("5c8ded8d39b4c70754ea3889")
}

func GetAllSession(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, err := dao.GetSessionByUserID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, session)
}

//func GetAllSession(w http.ResponseWriter, r *http.Request) {
//	sessions,err := dao.GetAllSession()
//	if err != nil {
//		respondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	respondWithJson(w, http.StatusOK, sessions)
//}

func RemoveSessions(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var session Session
	session, err := dao.GetSessionById(params["id"])
	if  err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := dao.RemoveUser(session.UserID.Hex()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := dao.RemoveSession(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_ = videoTimeDAO.RemoveBySessionId(session.SessionID.Hex())
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func InsertVideo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var video Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	video.VideoID = bson.NewObjectId()
	if err := videoDAO.Insert(video); err != nil {
		respondWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, video)

}
func GetAllVideo(w http.ResponseWriter, r *http.Request) {
	videos, err := videoDAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videos)
}
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := videoDAO.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
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
	if err := videoTimeDAO.Insert(videoTime); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videoTime)
}
func GetVideoTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoTime, err := videoTimeDAO.FindById(params["id"])
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
	if err := videoTimeDAO.Update(videoTime); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func GetAllVideoTime(w http.ResponseWriter, r *http.Request) {
	videoTimes, err := videoTimeDAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videoTimes)
}
func DeleteVideoTime(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := videoTimeDAO.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func InsertFeedback(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var fb = Feedback{}
	err := json.NewDecoder(r.Body).Decode(&fb)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = dao.InsertFeedback(fb)
	if err != nil {
		_, _ = fmt.Fprint(w, "Create fail")
	}
	_, _ = fmt.Fprintf(w, "Create successfully")
}

func GetAllFeedback(w http.ResponseWriter, r *http.Request)  {
	fbs, err := dao.GetAllFeedback()
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		_ = json.NewEncoder(w).Encode(fbs)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
	log.Print(code)
}
