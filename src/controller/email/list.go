package email

import (
	"encoding/json"
	"fmt"
	"net-http/myapp/controller"
	"net-http/myapp/repository"
	"net-http/myapp/utils"
	"net/http"
)

type EmailList struct {
}

func (receiver EmailList) ListEmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		response := controller.Response{Status: 405, Text: "Method Not Allowed"}
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}
	model := repository.EmailModel{}
	entities, err := model.GetAll()
	if err != nil {
		utils.WriteLogFile(err.Error())
		return
	}
	fmt.Println(entities)
	json.NewEncoder(w).Encode(entities)

}
