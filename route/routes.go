package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "welcome_robot/config"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Meta struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	ServerCode int    `json:"server_code"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/wr/v1",
		Index,
	},
	Route{
		"Create User",
		"POST",
		"/wr/v1/users",
		InsertUser,
	},
	Route{
		"Get User",
		"GET",
		"/wr/v1/users/{id}",
		GetUserById,
	},
	Route{
		"Remove User",
		"DELETE",
		"/wr/v1/users/{id}",
		RemoveUser,
	},
	Route{
		"Check duplicate name",
		"POST",
		"/wr/v1/checkname",
		CheckDuplicationName,
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
		"/wr/v1/users/{id}/session",
		GetSessionByUserId,
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
		"Delte Video",
		"DELETE",
		"/wr/v1/videos/{id}",
		DeleteVideo,
	},
	Route{
		"Get Video Time",
		"GET",
		"/wr/v1/videotimes/{id}",
		GetVideoTime,
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
		"Get All Feedback",
		"GET",
		"/wr/v1/feedbacks",
		GetAllFeedback,
	},
	Route{
		"Insert feedback",
		"POST",
		"/wr/v1/feedbacks",
		InsertFeedback,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_, _ = fmt.Fprintf(w, "Welcome Robot API")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	var data Response
	data.Meta.Status = "error"
	data.Meta.ServerCode = code
	data.Meta.Message = msg

	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
	log.Print(code)
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	var data Response
	data.Meta.Status = "success"
	data.Meta.ServerCode = code
	data.Data = payload

	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
	log.Print(code)
}
