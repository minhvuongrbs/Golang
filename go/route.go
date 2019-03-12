package _go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
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
		"GetAllSession",
		"GET",
		"/wr/v1/sessions",
		GetAllSession,
	},
	Route{
		"GetAllVideo",
		"GET",
		"/wr/v1/video",
		GetAllVideo,
	},
	Route{
		"Insert Video",
		"POST",
		"/wr/v1/video",
		InsertVideo,
	},
	Route{
		"Get Video Time",
		"GET",
		"/wr/v1/videotime/{id}",
		GetVideoTime,
	},
	Route{
		"Get all video time",
		"GET",
		"/wr/v1/videotime",
		GetAllVideoTime,
	},
	Route{
		"update video time",
		"PUT",
		"/wr/v1/videotime",
		UpdateVideoTime,
	},
	Route{
		"delete video time",
		"DELETE",
		"/wr/v1/videotime/{id}",
		DeleteVideoTime,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func GetAllSession(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All Sesstion")
}

func GetAllVideo(w http.ResponseWriter, r *http.Request) {
	videos, err := videoDAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, videos)
}
func InsertVideo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Insert Video")
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

func GetVideoTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoTime, err := videoTimeDAO.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
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
	log.Print("videotimeid"+videoTime.VideoTimeID)
	if  err := videoTimeDAO.Update(videoTime); err != nil {
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

func DeleteVideoTime(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := videoTimeDAO.Delete(params["id"]); err!=nil{
		respondWithError(w,http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w,http.StatusOK,map[string]string{"result":"success"})
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
