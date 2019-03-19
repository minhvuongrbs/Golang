package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"welcome_robot/dao"
	. "welcome_robot/models"
)

func InsertFeedback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var fb = Feedback{}
	err := json.NewDecoder(r.Body).Decode(&fb)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = dao.InsertFeedback(fb)
	log.Print(fb)
	if err != nil {
		_, _ = fmt.Fprint(w, "Create fail")
	}
	_, _ = fmt.Fprintf(w, "Create successfully")
}
func GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fbs, err := dao.FindAllFeedback()
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	} else {
		_ = json.NewEncoder(w).Encode(fbs)
	}
}
