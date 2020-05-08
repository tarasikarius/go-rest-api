package controllers

import (
	"encoding/json"
	"github.com/tarasikarius/go-rest-api/models"
	u "github.com/tarasikarius/go-rest-api/utils"
	"net/http"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))

		return
	}

	resp := models.Login(user.Email, user.Password)

	u.Respond(w, resp)
}

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request data"))

		return
	}

	resp := user.Create()

	u.Respond(w, resp)
}
